// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter_test

import (
	gc "gopkg.in/check.v1"
)

type cloudSpecUniterSuite struct {
}

var _ = gc.Suite(&cloudSpecUniterSuite{})

func (s *cloudSpecUniterSuite) TestStub(c *gc.C) {
	c.Skip(`This suite is missing tests for the following scenarios:
- TestGetCloudSpecReturnsSpecWhenTrusted: A test returning a correct cloud spec
for the given model when request is authorized.
- TestCloudAPIVersion: A test returning the cloud API version for the given 
model.
`)
}
