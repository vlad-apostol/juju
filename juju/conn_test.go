package juju_test

import (
	"io/ioutil"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/juju/environs/dummy"
	"launchpad.net/juju-core/juju/juju"
	"launchpad.net/juju-core/juju/testing"
	"os"
	"path/filepath"
	stdtesting "testing"
)

func Test(t *stdtesting.T) {
	srv := testing.StartZkServer()
	defer srv.Destroy()
	dummy.SetZookeeper(srv)
	defer dummy.SetZookeeper(nil)
	TestingT(t)
}

type ConnSuite struct{}

var _ = Suite(ConnSuite{})

func (ConnSuite) TestNewConn(c *C) {
	home := c.MkDir()
	defer os.Setenv("HOME", os.Getenv("HOME"))
	os.Setenv("HOME", home)
	conn, err := juju.NewConn("")
	c.Assert(conn, IsNil)
	c.Assert(err, ErrorMatches, ".*: no such file or directory")

	if err := os.Mkdir(filepath.Join(home, ".juju"), 0755); err != nil {
		c.Log("Could not create directory structure")
		c.Fail()
	}
	envs := filepath.Join(home, ".juju", "environments.yaml")
	err = ioutil.WriteFile(envs, []byte(`
default:
    erewhemos
environments:
    erewhemos:
        type: dummy
        zookeeper: true
`), 0644)
	if err != nil {
		c.Log("Could not create environments.yaml")
		c.Fail()
	}

	// Just run through a few operations on the dummy provider and verify that
	// they behave as expected.
	conn, err = juju.NewConn("")
	c.Assert(err, IsNil)
	defer conn.Close()
	st, err := conn.State()
	c.Assert(st, IsNil)
	c.Assert(err, ErrorMatches, "no state info available for this environ")
	err = conn.Bootstrap(false)
	c.Assert(err, IsNil)
	st, err = conn.State()
	c.Assert(err, IsNil)
	c.Assert(st, NotNil)
	err = conn.Destroy()
	c.Assert(err, IsNil)

	// Close the conn (thereby closing its state) a couple of times to
	// verify that multiple closes are safe.
	c.Assert(conn.Close(), IsNil)
	c.Assert(conn.Close(), IsNil)
}

func (ConnSuite) TestValidRegexps(c *C) {
	assertService := func(s string, expect bool) {
		c.Assert(juju.ValidService.MatchString(s), Equals, expect)
		c.Assert(juju.ValidUnit.MatchString(s+"/0"), Equals, expect)
		c.Assert(juju.ValidUnit.MatchString(s+"/99"), Equals, expect)
		c.Assert(juju.ValidUnit.MatchString(s+"/-1"), Equals, false)
		c.Assert(juju.ValidUnit.MatchString(s+"/blah"), Equals, false)
	}
	assertService("", false)
	assertService("33", false)
	assertService("wordpress", true)
	assertService("w0rd-pre55", true)
	assertService("foo2", true)
	assertService("foo-2", false)
	assertService("foo-2foo", true)
}
