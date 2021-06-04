package shardt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoreProcessor_Init_No_Remote(t *testing.T) {
	p := newCoreProcessor(3, 100, "self-addr", computeOptions())
	output := p.init()
	assert.Equal(t, []syncRemote(nil), output.sync)
}

func newStateWithPartitions(count int) State {
	partitions := make([]PartitionState, count)
	return State{
		Partitions: partitions,
	}
}

func TestCoreProcessor_Init_With_Remote(t *testing.T) {
	p := newCoreProcessor(3, 100, "self-addr",
		computeOptions(
			AddRemoteAddress("remote1"),
			AddRemoteAddress("remote2"),
		),
	)

	syncState := newStateWithPartitions(3)
	syncState.Nodes = map[string]NodeState{
		"self-addr": {
			Term:    1,
			Unix:    100,
			Version: 1,
			Status:  NodeStatusActive,
		},
	}

	output := p.init()
	assert.Equal(t, []syncRemote{
		{addr: "remote1", state: syncState},
		{addr: "remote2", state: syncState},
	}, output.sync)
}

func TestCoreProcessor_2_Remotes__Init_And_Finish_Sync(t *testing.T) {
	p := newCoreProcessor(3, 100, "self-addr",
		computeOptions(
			AddRemoteAddress("remote1"),
			AddRemoteAddress("remote2"),
		),
	)

	p.init()

	state := newStateWithPartitions(3)
	state.Nodes = map[string]NodeState{
		"self-addr": {
			Term:    1,
			Unix:    100,
			Version: 1,
			Status:  NodeStatusActive,
		},
		"remote1": {
			Term:    1,
			Unix:    200,
			Version: 1,
			Status:  NodeStatusActive,
		},
	}

	output := p.finishSync(true, state)
	assert.Equal(t, []syncRemote(nil), output.sync)
	assert.Equal(t, []bool{false, false, false}, output.expected)
}

func TestCoreProcessor_2_Remotes__Init_And_Finish_Sync_Second_Time(t *testing.T) {
	p := newCoreProcessor(3, 100, "self-addr",
		computeOptions(
			AddRemoteAddress("remote1"),
			AddRemoteAddress("remote2"),
		),
	)

	p.init()

	state := newStateWithPartitions(3)
	state.Nodes = map[string]NodeState{
		"self-addr": {
			Term:    1,
			Unix:    100,
			Version: 1,
			Status:  NodeStatusActive,
		},
		"remote1": {
			Term:    1,
			Unix:    200,
			Version: 1,
			Status:  NodeStatusActive,
		},
	}

	p.finishSync(true, state)
	output := p.finishSync(true, state)

	assert.Equal(t, []bool{true, true, false}, output)
}
