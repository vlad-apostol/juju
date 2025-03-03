// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package status_test

import (
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/status"
)

type viableSuite struct{}

var _ = gc.Suite(&viableSuite{})

func (s *viableSuite) TestIsMachineViable(c *gc.C) {
	testcases := []struct {
		name   string
		status status.StatusInfo
		viable bool
	}{
		{
			name: "started",
			status: status.StatusInfo{
				Status: status.Started,
			},
			viable: true,
		},
		{
			name: "pending",
			status: status.StatusInfo{
				Status: status.Pending,
			},
			viable: false,
		},
		{
			name: "down",
			status: status.StatusInfo{
				Status: status.Down,
			},
			viable: false,
		},
		{
			name: "stopped",
			status: status.StatusInfo{
				Status: status.Stopped,
			},
			viable: false,
		},
		{
			name: "error",
			status: status.StatusInfo{
				Status: status.Error,
			},
			viable: false,
		},
		{
			name: "unknown",
			status: status.StatusInfo{
				Status: status.Unknown,
			},
			viable: false,
		},
	}
	for _, v := range testcases {
		c.Logf("Testing machine status %s", v.name)
		c.Assert(status.IsMachineViable(v.status), gc.Equals, v.viable)
	}
}

func (s *viableSuite) TestIsInstanceViable(c *gc.C) {
	testcases := []struct {
		name   string
		status status.StatusInfo
		viable bool
	}{
		{
			name: "running",
			status: status.StatusInfo{
				Status: status.Running,
			},
			viable: true,
		},
		{
			name: "empty",
			status: status.StatusInfo{
				Status: status.Empty,
			},
			viable: false,
		},
		{
			name: "allocating",
			status: status.StatusInfo{
				Status: status.Allocating,
			},
			viable: false,
		},
		{
			name: "error",
			status: status.StatusInfo{
				Status: status.Error,
			},
			viable: false,
		},
		{
			name: "provisioning error",
			status: status.StatusInfo{
				Status: status.ProvisioningError,
			},
			viable: false,
		},
		{
			name: "unknown",
			status: status.StatusInfo{
				Status: status.Unknown,
			},
			viable: false,
		},
	}
	for _, v := range testcases {
		c.Logf("Testing instance status %s", v.name)
		c.Assert(status.IsInstanceViable(v.status), gc.Equals, v.viable)
	}
}

func (s *viableSuite) TestIsAgentViable(c *gc.C) {
	testcases := []struct {
		name   string
		status status.StatusInfo
		viable bool
	}{
		{
			name: "idle",
			status: status.StatusInfo{
				Status: status.Idle,
			},
			viable: true,
		},
		{
			name: "executing",
			status: status.StatusInfo{
				Status: status.Executing,
			},
			viable: true,
		},
		{
			name: "allocating",
			status: status.StatusInfo{
				Status: status.Allocating,
			},
			viable: false,
		},
		{
			name: "error",
			status: status.StatusInfo{
				Status: status.Error,
			},
			viable: false,
		},
		{
			name: "failed",
			status: status.StatusInfo{
				Status: status.Failed,
			},
			viable: false,
		},
		{
			name: "rebooting",
			status: status.StatusInfo{
				Status: status.Rebooting,
			},
			viable: false,
		},
		{
			name: "unknown",
			status: status.StatusInfo{
				Status: status.Unknown,
			},
			viable: false,
		},
	}
	for _, v := range testcases {
		c.Logf("Testing agent status %s", v.name)
		c.Assert(status.IsAgentViable(v.status), gc.Equals, v.viable)
	}
}

func (s *viableSuite) TestIsUnitWorkloadViable(c *gc.C) {
	testcases := []struct {
		name   string
		status status.StatusInfo
		viable bool
	}{
		{
			name: "active",
			status: status.StatusInfo{
				Status: status.Active,
			},
			viable: true,
		},
		{
			name: "maintenance installing charm",
			status: status.StatusInfo{
				Status:  status.Maintenance,
				Message: status.MessageInstallingCharm,
			},
			viable: false,
		},
		{
			name: "maintenance",
			status: status.StatusInfo{
				Status: status.Maintenance,
			},
			viable: true,
		},
		{
			name: "waiting for machine",
			status: status.StatusInfo{
				Status:  status.Waiting,
				Message: status.MessageWaitForMachine,
			},
			viable: false,
		},
		{
			name: "waiting installing agent",
			status: status.StatusInfo{
				Status:  status.Waiting,
				Message: status.MessageInstallingAgent,
			},
			viable: false,
		},
		{
			name: "waiting initializing agent",
			status: status.StatusInfo{
				Status:  status.Waiting,
				Message: status.MessageInitializingAgent,
			},
			viable: false,
		},
		{
			name: "waiting",
			status: status.StatusInfo{
				Status: status.Waiting,
			},
			viable: true,
		},
		{
			name: "blocked",
			status: status.StatusInfo{
				Status: status.Blocked,
			},
			viable: false,
		},
		{
			name: "error",
			status: status.StatusInfo{
				Status: status.Error,
			},
			viable: false,
		},
		{
			name: "terminated",
			status: status.StatusInfo{
				Status: status.Terminated,
			},
			viable: false,
		},
		{
			name: "unknown",
			status: status.StatusInfo{
				Status: status.Unknown,
			},
			viable: false,
		},
	}
	for _, v := range testcases {
		c.Logf("Testing unit workload status %s", v.name)
		c.Assert(status.IsUnitWorkloadViable(v.status), gc.Equals, v.viable)
	}
}
