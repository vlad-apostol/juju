package constraints

import (
	"github.com/juju/testing"
	gc "gopkg.in/check.v1"
)

type archSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&archSuite{})

func (s archSuite) TestConstraintArch(c *gc.C) {
	a := ConstraintArch(MustParse("mem=4G"), nil)
	c.Assert(a, gc.Equals, "amd64")
	a = ConstraintArch(MustParse("arch=arm64"), nil)
	c.Assert(a, gc.Equals, "arm64")
	defaultCons := MustParse("arch=arm64")
	a = ConstraintArch(MustParse("mem=4G"), &defaultCons)
	c.Assert(a, gc.Equals, "arm64")
	a = ConstraintArch(MustParse("arch=s390x"), &defaultCons)
	c.Assert(a, gc.Equals, "s390x")
}
