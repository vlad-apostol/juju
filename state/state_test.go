package state_test

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/charm"
	"launchpad.net/juju-core/environs/config"
	"launchpad.net/juju-core/state"
	"launchpad.net/juju-core/testing"
	"net/url"
	"sort"
	"time"
)

type D []bson.DocElem

type StateSuite struct {
	ConnSuite
}

var _ = Suite(&StateSuite{})

func (s *StateSuite) TestDialAgain(c *C) {
	// Ensure idempotent operations on Dial are working fine.
	for i := 0; i < 2; i++ {
		st, err := state.Open(&state.Info{Addrs: []string{testing.MgoAddr}})
		c.Assert(err, IsNil)
		c.Assert(st.Close(), IsNil)
	}
}

func (s *StateSuite) TestIsNotFound(c *C) {
	err1 := fmt.Errorf("unrelated error")
	err2 := &state.NotFoundError{}
	c.Assert(state.IsNotFound(err1), Equals, false)
	c.Assert(state.IsNotFound(err2), Equals, true)
}

func (s *StateSuite) TestAddCharm(c *C) {
	// Check that adding charms from scratch works correctly.
	ch := testing.Charms.Dir("series", "dummy")
	curl := charm.MustParseURL(
		fmt.Sprintf("local:series/%s-%d", ch.Meta().Name, ch.Revision()),
	)
	bundleURL, err := url.Parse("http://bundles.example.com/dummy-1")
	c.Assert(err, IsNil)
	dummy, err := s.State.AddCharm(ch, curl, bundleURL, "dummy-1-sha256")
	c.Assert(err, IsNil)
	c.Assert(dummy.URL().String(), Equals, curl.String())

	doc := state.CharmDoc{}
	err = s.charms.FindId(curl).One(&doc)
	c.Assert(err, IsNil)
	c.Logf("%#v", doc)
	c.Assert(doc.URL, DeepEquals, curl)
}

func (s *StateSuite) AssertMachineCount(c *C, expect int) {
	ms, err := s.State.AllMachines()
	c.Assert(err, IsNil)
	c.Assert(len(ms), Equals, expect)
}

func (s *StateSuite) TestAddMachine(c *C) {
	machine0, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	c.Assert(machine0.Id(), Equals, 0)
	machine1, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	c.Assert(machine1.Id(), Equals, 1)

	machines := s.AllMachines(c)
	c.Assert(machines, DeepEquals, []int{0, 1})
}

func (s *StateSuite) TestRemoveMachine(c *C) {
	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	_, err = s.State.AddMachine()
	c.Assert(err, IsNil)
	err = s.State.RemoveMachine(machine.Id())
	c.Assert(err, ErrorMatches, "cannot remove machine 0: machine is not dead")
	err = machine.EnsureDead()
	c.Assert(err, IsNil)
	err = s.State.RemoveMachine(machine.Id())
	c.Assert(err, IsNil)

	machines := s.AllMachines(c)
	c.Assert(machines, DeepEquals, []int{1})

	// Removing a non-existing machine has to fail.
	// BUG(aram): use error strings from state.
	err = s.State.RemoveMachine(machine.Id())
	c.Assert(err, ErrorMatches, "cannot remove machine 0: .*")
}

func (s *StateSuite) TestReadMachine(c *C) {
	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	expectedId := machine.Id()
	machine, err = s.State.Machine(expectedId)
	c.Assert(err, IsNil)
	c.Assert(machine.Id(), Equals, expectedId)
}

func (s *StateSuite) TestMachineNotFound(c *C) {
	_, err := s.State.Machine(0)
	c.Assert(err, ErrorMatches, "machine 0 not found")
	c.Assert(state.IsNotFound(err), Equals, true)
}

func (s *StateSuite) TestAllMachines(c *C) {
	c.Skip("Marshalling of agent tools is currently broken")
	numInserts := 42
	for i := 0; i < numInserts; i++ {
		m, err := s.State.AddMachine()
		c.Assert(err, IsNil)
		err = m.SetInstanceId(fmt.Sprintf("foo-%d", i))
		c.Assert(err, IsNil)
		err = m.SetAgentTools(newTools("7.8.9-foo-bar", "http://arble.tgz"))
		c.Assert(err, IsNil)
		err = m.EnsureDying()
		c.Assert(err, IsNil)
	}
	s.AssertMachineCount(c, numInserts)
	ms, _ := s.State.AllMachines()
	for i, m := range ms {
		c.Assert(m.Id(), Equals, i)
		instId, err := m.InstanceId()
		c.Assert(err, IsNil)
		c.Assert(instId, Equals, fmt.Sprintf("foo-%d", i))
		tools, err := m.AgentTools()
		c.Check(err, IsNil)
		c.Check(tools, DeepEquals, newTools("7.8.9-foo-bar", "http://arble.tgz"))
		c.Assert(m.Life(), Equals, state.Dying)
	}
}

func (s *StateSuite) TestAddService(c *C) {
	charm := s.AddTestingCharm(c, "dummy")
	_, err := s.State.AddService("haha/borken", charm)
	c.Assert(err, ErrorMatches, `"haha/borken" is not a valid service name`)
	_, err = s.State.Service("haha/borken")
	c.Assert(err, ErrorMatches, `"haha/borken" is not a valid service name`)

	wordpress, err := s.State.AddService("wordpress", charm)
	c.Assert(err, IsNil)
	c.Assert(wordpress.Name(), Equals, "wordpress")
	mysql, err := s.State.AddService("mysql", charm)
	c.Assert(err, IsNil)
	c.Assert(mysql.Name(), Equals, "mysql")

	// Check that retrieving the new created services works correctly.
	wordpress, err = s.State.Service("wordpress")
	c.Assert(err, IsNil)
	c.Assert(wordpress.Name(), Equals, "wordpress")
	ch, _, err := wordpress.Charm()
	c.Assert(err, IsNil)
	c.Assert(ch.URL(), DeepEquals, charm.URL())
	mysql, err = s.State.Service("mysql")
	c.Assert(err, IsNil)
	c.Assert(mysql.Name(), Equals, "mysql")
	ch, _, err = mysql.Charm()
	c.Assert(err, IsNil)
	c.Assert(ch.URL(), DeepEquals, charm.URL())
}

func (s *StateSuite) TestServiceNotFound(c *C) {
	_, err := s.State.Service("bummer")
	c.Assert(err, ErrorMatches, `service "bummer" not found`)
	c.Assert(state.IsNotFound(err), Equals, true)
}

func (s *StateSuite) TestRemoveService(c *C) {
	charm := s.AddTestingCharm(c, "dummy")
	service, err := s.State.AddService("wordpress", charm)
	c.Assert(err, IsNil)

	// Remove of existing service.
	err = s.State.RemoveService(service)
	c.Assert(err, ErrorMatches, `cannot remove service "wordpress": service is not dead`)
	err = service.EnsureDead()
	c.Assert(err, IsNil)
	err = s.State.RemoveService(service)
	c.Assert(err, IsNil)
	_, err = s.State.Service("wordpress")
	c.Assert(err, ErrorMatches, `service "wordpress" not found`)

	err = s.State.RemoveService(service)
	c.Assert(err, IsNil)
}

func (s *StateSuite) TestAllServices(c *C) {
	charm := s.AddTestingCharm(c, "dummy")
	services, err := s.State.AllServices()
	c.Assert(err, IsNil)
	c.Assert(len(services), Equals, 0)

	// Check that after adding services the result is ok.
	_, err = s.State.AddService("wordpress", charm)
	c.Assert(err, IsNil)
	services, err = s.State.AllServices()
	c.Assert(err, IsNil)
	c.Assert(len(services), Equals, 1)

	_, err = s.State.AddService("mysql", charm)
	c.Assert(err, IsNil)
	services, err = s.State.AllServices()
	c.Assert(err, IsNil)
	c.Assert(len(services), Equals, 2)

	// Check the returned service, order is defined by sorted keys.
	c.Assert(services[0].Name(), Equals, "wordpress")
	c.Assert(services[1].Name(), Equals, "mysql")
}

func (s *StateSuite) TestEnvironConfig(c *C) {
	initial := map[string]interface{}{
		"name":            "test",
		"type":            "test",
		"authorized-keys": "i-am-a-key",
		"default-series":  "precise",
		"development":     true,
	}
	env, err := config.New(initial)
	c.Assert(err, IsNil)
	err = s.State.SetEnvironConfig(env)
	c.Assert(err, IsNil)
	env, err = s.State.EnvironConfig()
	c.Assert(err, IsNil)
	current := env.AllAttrs()
	c.Assert(current, DeepEquals, initial)

	current["authorized-keys"] = "i-am-a-new-key"
	env, err = config.New(current)
	c.Assert(err, IsNil)
	err = s.State.SetEnvironConfig(env)
	c.Assert(err, IsNil)
	env, err = s.State.EnvironConfig()
	c.Assert(err, IsNil)
	final := env.AllAttrs()
	c.Assert(final, DeepEquals, current)
}

var machinesWatchTests = []struct {
	test  func(*C, *state.State)
	alive []int
	dead  []int
}{
	{
		test: func(_ *C, _ *state.State) {},
	}, {
		test: func(c *C, s *state.State) {
			_, err := s.AddMachine()
			c.Assert(err, IsNil)
		},
		alive: []int{0},
	}, {
		test: func(c *C, s *state.State) {
			_, err := s.AddMachine()
			c.Assert(err, IsNil)
			m0, err := s.Machine(0)
			c.Assert(err, IsNil)
			err = m0.SetInstanceId("spam")
			c.Assert(err, IsNil)
		},
		alive: []int{1},
	}, {
		test: func(c *C, s *state.State) {
			_, err := s.AddMachine()
			c.Assert(err, IsNil)
			_, err = s.AddMachine()
			c.Assert(err, IsNil)
		},
		alive: []int{2, 3},
	}, {
		test: func(c *C, s *state.State) {
			m3, err := s.Machine(3)
			c.Assert(err, IsNil)
			err = m3.EnsureDead()
			c.Assert(err, IsNil)
		},
		dead: []int{3},
	}, {
		test: func(c *C, s *state.State) {
			m0, err := s.Machine(0)
			c.Assert(err, IsNil)
			err = m0.EnsureDead()
			c.Assert(err, IsNil)
			m2, err := s.Machine(2)
			c.Assert(err, IsNil)
			err = m2.EnsureDead()
			c.Assert(err, IsNil)
		},
		dead: []int{0, 2},
	}, {
		test: func(c *C, s *state.State) {
			_, err := s.AddMachine()
			c.Assert(err, IsNil)
			m1, err := s.Machine(1)
			c.Assert(err, IsNil)
			err = m1.EnsureDead()
			c.Assert(err, IsNil)
		},
		alive: []int{4},
		dead:  []int{1},
	}, {
		test: func(c *C, s *state.State) {
			machines := [20]*state.Machine{}
			var err error
			for i := 0; i < len(machines); i++ {
				machines[i], err = s.AddMachine()
				c.Assert(err, IsNil)
			}
			for i := 0; i < len(machines); i++ {
				err = machines[i].SetInstanceId("spam" + fmt.Sprint(i))
				c.Assert(err, IsNil)
			}
			for i := 10; i < len(machines); i++ {
				err = machines[i].EnsureDead()
				c.Assert(err, IsNil)
			}
		},
		alive: []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
	}, {
		test: func(c *C, s *state.State) {
			_, err := s.AddMachine()
			c.Assert(err, IsNil)
			m9, err := s.Machine(9)
			c.Assert(err, IsNil)
			err = m9.EnsureDead()
			c.Assert(err, IsNil)
		},
		alive: []int{25},
		dead:  []int{9},
	},
}

func (s *StateSuite) TestWatchMachines(c *C) {
	machineWatcher := s.State.WatchMachines()
	defer func() {
		c.Assert(machineWatcher.Stop(), IsNil)
	}()
	for i, test := range machinesWatchTests {
		c.Logf("test %d", i)
		test.test(c, s.State)
		s.State.StartSync()
		got := &state.MachinesChange{}
		for {
			select {
			case new, ok := <-machineWatcher.Changes():
				c.Assert(ok, Equals, true)
				addMachinesChange(got, new)
				if moreMachinesRequired(got, test.alive, test.dead) {
					continue
				}
				assertSameMachines(c, got, test.alive, test.dead)
			case <-time.After(500 * time.Millisecond):
				c.Fatalf("did not get change, want: alive: %#v, dead: %#v, got: %#v", test.alive, test.dead, got)
			}
			break
		}
	}
	select {
	case got := <-machineWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(100 * time.Millisecond):
	}
}

func moreMachinesRequired(got *state.MachinesChange, alive, dead []int) bool {
	return len(got.Alive)+len(got.Dead) < len(alive)+len(dead)
}

func addMachinesChange(change *state.MachinesChange, more *state.MachinesChange) {
	change.Alive = append(change.Alive, more.Alive...)
	change.Dead = append(change.Dead, more.Dead...)
}

func assertSameMachines(c *C, got *state.MachinesChange, alive, dead []int) {
	c.Assert(got, NotNil)
	sort.Ints(got.Alive)
	sort.Ints(got.Dead)
	c.Assert(got.Alive, DeepEquals, alive)
	c.Assert(got.Dead, DeepEquals, dead)
}

var servicesWatchTests = []struct {
	test    func(*C, *state.State, *state.Charm)
	added   []string
	removed []string
}{
	{
		test:  func(_ *C, _ *state.State, _ *state.Charm) {},
		added: []string{},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			_, err := s.AddService("s0", ch)
			c.Assert(err, IsNil)
		},
		added: []string{"s0"},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			_, err := s.AddService("s1", ch)
			c.Assert(err, IsNil)
		},
		added: []string{"s1"},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			_, err := s.AddService("s2", ch)
			c.Assert(err, IsNil)
			_, err = s.AddService("s3", ch)
			c.Assert(err, IsNil)
		},
		added: []string{"s2", "s3"},
	},
	{
		test: func(c *C, s *state.State, _ *state.Charm) {
			svc3, err := s.Service("s3")
			c.Assert(err, IsNil)
			err = svc3.EnsureDead()
			c.Assert(err, IsNil)
			err = s.RemoveService(svc3)
			c.Assert(err, IsNil)
		},
		removed: []string{"s3"},
	},
	{
		test: func(c *C, s *state.State, _ *state.Charm) {
			svc0, err := s.Service("s0")
			c.Assert(err, IsNil)
			err = svc0.EnsureDead()
			c.Assert(err, IsNil)
			err = s.RemoveService(svc0)
			c.Assert(err, IsNil)
			svc2, err := s.Service("s2")
			c.Assert(err, IsNil)
			err = svc2.EnsureDead()
			c.Assert(err, IsNil)
			err = s.RemoveService(svc2)
			c.Assert(err, IsNil)
		},
		removed: []string{"s0", "s2"},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			_, err := s.AddService("s4", ch)
			c.Assert(err, IsNil)
			svc1, err := s.Service("s1")
			c.Assert(err, IsNil)
			err = svc1.EnsureDead()
			c.Assert(err, IsNil)
			err = s.RemoveService(svc1)
			c.Assert(err, IsNil)
		},
		added:   []string{"s4"},
		removed: []string{"s1"},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			services := [20]*state.Service{}
			var err error
			for i := 0; i < len(services); i++ {
				services[i], err = s.AddService("ss"+fmt.Sprint(i), ch)
				c.Assert(err, IsNil)
			}
			for i := 10; i < len(services); i++ {
				err = services[i].EnsureDead()
				c.Assert(err, IsNil)
				err = s.RemoveService(services[i])
				c.Assert(err, IsNil)
			}
		},
		added: []string{"ss0", "ss1", "ss2", "ss3", "ss4", "ss5", "ss6", "ss7", "ss8", "ss9"},
	},
	{
		test: func(c *C, s *state.State, ch *state.Charm) {
			_, err := s.AddService("twenty-five", ch)
			c.Assert(err, IsNil)
			svc9, err := s.Service("ss9")
			c.Assert(err, IsNil)
			err = svc9.EnsureDead()
			c.Assert(err, IsNil)
			err = s.RemoveService(svc9)
			c.Assert(err, IsNil)
		},
		added:   []string{"twenty-five"},
		removed: []string{"ss9"},
	},
}

func (s *StateSuite) TestWatchServices(c *C) {
	serviceWatcher := s.State.WatchServices()
	defer func() {
		c.Assert(serviceWatcher.Stop(), IsNil)
	}()
	charm := s.AddTestingCharm(c, "dummy")
	for i, test := range servicesWatchTests {
		c.Logf("test %d", i)
		test.test(c, s.State, charm)
		s.State.StartSync()
		got := &state.ServicesChange{}
		for {
			select {
			case new, ok := <-serviceWatcher.Changes():
				c.Assert(ok, Equals, true)
				addServiceChanges(got, new)
				if moreServicesRequired(got, test.added, test.removed) {
					continue
				}
				assertSameServices(c, got, test.added, test.removed)
			case <-time.After(500 * time.Millisecond):
				c.Fatalf("did not get change, want: added: %#v, removed: %#v, got: %#v", test.added, test.removed, got)
			}
			break
		}
	}
	select {
	case got := <-serviceWatcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(100 * time.Millisecond):
	}
}

func moreServicesRequired(got *state.ServicesChange, added, removed []string) bool {
	return len(got.Added)+len(got.Removed) < len(added)+len(removed)
}

func addServiceChanges(changes *state.ServicesChange, more *state.ServicesChange) {
	changes.Added = append(changes.Added, more.Added...)
	changes.Removed = append(changes.Removed, more.Removed...)
}

type serviceSlice []*state.Service

func (m serviceSlice) Len() int           { return len(m) }
func (m serviceSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m serviceSlice) Less(i, j int) bool { return m[i].Name() < m[j].Name() }

func assertSameServices(c *C, change *state.ServicesChange, added, removed []string) {
	c.Assert(change, NotNil)
	if len(added) == 0 {
		added = nil
	}
	if len(removed) == 0 {
		removed = nil
	}
	sort.Sort(serviceSlice(change.Added))
	sort.Sort(serviceSlice(change.Removed))
	var got []string
	for _, g := range change.Added {
		got = append(got, g.Name())
	}
	c.Assert(got, DeepEquals, added)
	got = nil
	for _, g := range change.Removed {
		got = append(got, g.Name())
	}
	c.Assert(got, DeepEquals, removed)
}

func (s *StateSuite) TestInitialize(c *C) {
	m := map[string]interface{}{
		"type":            "dummy",
		"name":            "lisboa",
		"authorized-keys": "i-am-a-key",
		"default-series":  "precise",
		"development":     true,
	}
	cfg, err := config.New(m)
	c.Assert(err, IsNil)
	st, err := state.Initialize(s.StateInfo(c), cfg)
	c.Assert(err, IsNil)
	c.Assert(st, NotNil)
	defer st.Close()
	env, err := st.EnvironConfig()
	c.Assert(env.AllAttrs(), DeepEquals, m)
}

func (s *StateSuite) TestDoubleInitialize(c *C) {
	m := map[string]interface{}{
		"type":            "dummy",
		"name":            "lisboa",
		"authorized-keys": "i-am-a-key",
		"default-series":  "precise",
		"development":     true,
	}
	cfg, err := config.New(m)
	c.Assert(err, IsNil)
	st, err := state.Initialize(s.StateInfo(c), cfg)
	c.Assert(err, IsNil)
	c.Assert(st, NotNil)
	env1, err := st.EnvironConfig()
	st.Close()

	// initialize again, there should be no error and the 
	// environ config should not change.
	m = map[string]interface{}{
		"type":            "dummy",
		"name":            "sydney",
		"authorized-keys": "i-am-not-an-animal",
		"default-series":  "xanadu",
		"development":     false,
	}
	cfg, err = config.New(m)
	c.Assert(err, IsNil)
	st, err = state.Initialize(s.StateInfo(c), cfg)
	c.Assert(err, IsNil)
	c.Assert(st, NotNil)
	env2, err := st.EnvironConfig()
	st.Close()

	c.Assert(env1.AllAttrs(), DeepEquals, env2.AllAttrs())
}

var sortPortsTests = []struct {
	have, want []state.Port
}{
	{nil, []state.Port{}},
	{[]state.Port{{"b", 1}, {"a", 99}, {"a", 1}}, []state.Port{{"a", 1}, {"a", 99}, {"b", 1}}},
}

func (*StateSuite) TestSortPorts(c *C) {
	for _, t := range sortPortsTests {
		p := make([]state.Port, len(t.have))
		copy(p, t.have)
		state.SortPorts(p)
		c.Check(p, DeepEquals, t.want)
		state.SortPorts(p)
		c.Check(p, DeepEquals, t.want)
	}
}

func (*StateSuite) TestNameChecks(c *C) {
	assertService := func(s string, expect bool) {
		c.Assert(state.IsServiceName(s), Equals, expect)
		c.Assert(state.IsUnitName(s+"/0"), Equals, expect)
		c.Assert(state.IsUnitName(s+"/99"), Equals, expect)
		c.Assert(state.IsUnitName(s+"/-1"), Equals, false)
		c.Assert(state.IsUnitName(s+"/blah"), Equals, false)
	}
	assertService("", false)
	assertService("33", false)
	assertService("wordpress", true)
	assertService("w0rd-pre55", true)
	assertService("foo2", true)
	assertService("foo-2", false)
	assertService("foo-2foo", true)
}

type attrs map[string]interface{}

var watchEnvironConfigTests = []attrs{
	{
		"type":            "my-type",
		"name":            "my-name",
		"authorized-keys": "i-am-a-key",
	},
	{
		// Add an attribute.
		"type":            "my-type",
		"name":            "my-name",
		"default-series":  "my-series",
		"authorized-keys": "i-am-a-key",
	},
	{
		// Set a new attribute value.
		"type":            "my-type",
		"name":            "my-new-name",
		"default-series":  "my-series",
		"authorized-keys": "i-am-a-key",
	},
}

func (s *StateSuite) TestWatchEnvironConfig(c *C) {
	watcher := s.State.WatchEnvironConfig()
	defer func() {
		c.Assert(watcher.Stop(), IsNil)
	}()
	for i, test := range watchEnvironConfigTests {
		c.Logf("test %d", i)
		change, err := config.New(test)
		c.Assert(err, IsNil)
		err = s.State.SetEnvironConfig(change)
		c.Assert(err, IsNil)
		s.State.StartSync()
		select {
		case got, ok := <-watcher.Changes():
			c.Assert(ok, Equals, true)
			c.Assert(got.AllAttrs(), DeepEquals, change.AllAttrs())
		case <-time.After(500 * time.Millisecond):
			c.Fatalf("did not get change: %#v", test)
		}
	}

	select {
	case got := <-watcher.Changes():
		c.Fatalf("got unexpected change: %#v", got)
	case <-time.After(50 * time.Millisecond):
	}
}

func (s *StateSuite) TestWatchEnvironConfigInitialConfig(c *C) {
	cfg, err := config.New(watchEnvironConfigTests[0])
	c.Assert(err, IsNil)
	err = s.State.SetEnvironConfig(cfg)
	c.Assert(err, IsNil)
	s.State.Sync()
	watcher := s.State.WatchEnvironConfig()
	defer watcher.Stop()
	select {
	case got, ok := <-watcher.Changes():
		c.Assert(ok, Equals, true)
		c.Assert(got.AllAttrs(), DeepEquals, cfg.AllAttrs())
	case <-time.After(500 * time.Millisecond):
		c.Fatalf("did not get change")
	}
}

func (s *StateSuite) TestWatchEnvironConfigInvalidConfig(c *C) {
	m := map[string]interface{}{
		"type":            "dummy",
		"name":            "lisboa",
		"authorized-keys": "i-am-a-key",
	}
	cfg1, err := config.New(m)
	c.Assert(err, IsNil)
	err = s.State.SetEnvironConfig(cfg1)
	c.Assert(err, IsNil)

	// Corrupt the environment configuration.
	settings := s.Session.DB("juju").C("settings")
	err = settings.UpdateId("e", bson.D{{"$unset", bson.D{{"name", 1}}}})
	c.Assert(err, IsNil)

	s.State.Sync()

	// Start watching the configuration.
	watcher := s.State.WatchEnvironConfig()
	defer watcher.Stop()
	done := make(chan *config.Config)
	go func() {
		select {
		case cfg, ok := <-watcher.Changes():
			if !ok {
				c.Errorf("watcher channel closed")
			} else {
				done <- cfg
			}
		case <-time.After(5 * time.Second):
			c.Fatalf("no environment configuration observed")
		}
	}()

	s.State.Sync()

	// The invalid configuration must not have been generated.
	select {
	case <-done:
		c.Fatalf("configuration returned too soon")
	case <-time.After(100 * time.Millisecond):
	}

	// Fix the configuration.
	cfg2, err := config.New(map[string]interface{}{
		"type":            "dummy",
		"name":            "lisboa",
		"authorized-keys": "new-key",
	})
	c.Assert(err, IsNil)
	err = s.State.SetEnvironConfig(cfg2)
	c.Assert(err, IsNil)

	s.State.StartSync()
	select {
	case cfg3 := <-done:
		c.Assert(cfg3.AllAttrs(), DeepEquals, cfg2.AllAttrs())
	case <-time.After(5 * time.Second):
		c.Fatalf("no environment configuration observed")
	}
}
