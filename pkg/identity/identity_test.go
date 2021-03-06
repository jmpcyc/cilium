// Copyright 2016-2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !privileged_tests

package identity

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type IdentityTestSuite struct{}

var _ = Suite(&IdentityTestSuite{})

func (s *IdentityTestSuite) TestReservedID(c *C) {
	i := GetReservedID("host")
	c.Assert(i, Equals, NumericIdentity(1))
	c.Assert(i.String(), Equals, "host")

	i = GetReservedID("world")
	c.Assert(i, Equals, NumericIdentity(2))
	c.Assert(i.String(), Equals, "world")

	// This is an obsoleted identity, we verify that it returns 0
	i = GetReservedID("cluster")
	c.Assert(i, Equals, NumericIdentity(0))
	c.Assert(i.String(), Equals, "0")

	i = GetReservedID("health")
	c.Assert(i, Equals, NumericIdentity(4))
	c.Assert(i.String(), Equals, "health")

	i = GetReservedID("init")
	c.Assert(i, Equals, NumericIdentity(5))
	c.Assert(i.String(), Equals, "init")

	i = GetReservedID("unmanaged")
	c.Assert(i, Equals, NumericIdentity(3))
	c.Assert(i.String(), Equals, "unmanaged")

	c.Assert(GetReservedID("unknown"), Equals, IdentityUnknown)
	unknown := NumericIdentity(700)
	c.Assert(unknown.String(), Equals, "700")
}

func (s *IdentityTestSuite) TestIsReservedIdentity(c *C) {
	c.Assert(ReservedIdentityHealth.IsReservedIdentity(), Equals, true)
	c.Assert(ReservedIdentityHost.IsReservedIdentity(), Equals, true)
	c.Assert(ReservedIdentityWorld.IsReservedIdentity(), Equals, true)
	c.Assert(ReservedIdentityInit.IsReservedIdentity(), Equals, true)
	c.Assert(ReservedIdentityUnmanaged.IsReservedIdentity(), Equals, true)

	c.Assert(NumericIdentity(123456).IsReservedIdentity(), Equals, false)
}
