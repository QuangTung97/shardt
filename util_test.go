package shardt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeStateLess(t *testing.T) {
	table := []struct {
		name     string
		a        NodeState
		b        NodeState
		expected bool
	}{
		{
			name:     "both-empty",
			expected: false,
		},
		{
			name: "a-term-less",
			a: NodeState{
				Term: 10,
			},
			b: NodeState{
				Term: 11,
			},
			expected: true,
		},
		{
			name: "a-term-greater",
			a: NodeState{
				Term: 11,
				Unix: 80,
			},
			b: NodeState{
				Term: 10,
				Unix: 100,
			},
			expected: false,
		},
		{
			name: "term-equal",
			a: NodeState{
				Term: 10,
			},
			b: NodeState{
				Term: 10,
			},
			expected: false,
		},
		{
			name: "a-unix-less",
			a: NodeState{
				Term: 11,
				Unix: 100,
			},
			b: NodeState{
				Term: 11,
				Unix: 101,
			},
			expected: true,
		},
		{
			name: "a-unix-greater",
			a: NodeState{
				Term:    11,
				Unix:    101,
				Version: 20,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 21,
			},
			expected: false,
		},
		{
			name: "unix-equal",
			a: NodeState{
				Term: 11,
				Unix: 100,
			},
			b: NodeState{
				Term: 11,
				Unix: 100,
			},
			expected: false,
		},
		{
			name: "a-version-less",
			a: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 21,
			},
			expected: true,
		},
		{
			name: "a-version-greater",
			a: NodeState{
				Term:    11,
				Unix:    100,
				Version: 21,
				Status:  NodeStatusActive,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusOutOfSync,
			},
			expected: false,
		},
		{
			name: "version-equal",
			a: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
			},
			expected: false,
		},
		{
			name: "a-status-less",
			a: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusActive,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusOutOfSync,
			},
			expected: true,
		},
		{
			name: "a-status-greater",
			a: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
			b: NodeState{
				Term:    11,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusOutOfSync,
			},
			expected: false,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {

			result := nodeStateLess(e.a, e.b)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestPartitionStateLess(t *testing.T) {
	table := []struct {
		name     string
		a        PartitionState
		b        PartitionState
		expected bool
	}{
		{
			name:     "both-empty",
			expected: false,
		},
		{
			name: "a-term-less",
			a: PartitionState{
				Term: 10,
			},
			b: PartitionState{
				Term: 11,
			},
			expected: true,
		},
		{
			name: "a-term-greater",
			a: PartitionState{
				Term: 11,
				Addr: "addr1",
			},
			b: PartitionState{
				Term: 10,
				Addr: "addr2",
			},
			expected: false,
		},
		{
			name: "term-equal",
			a: PartitionState{
				Term: 11,
			},
			b: PartitionState{
				Term: 11,
			},
			expected: false,
		},
		{
			name: "a-addr-less",
			a: PartitionState{
				Term: 11,
				Addr: "addr1",
			},
			b: PartitionState{
				Term: 11,
				Addr: "addr2",
			},
			expected: true,
		},
		{
			name: "a-addr-greater",
			a: PartitionState{
				Term:   11,
				Addr:   "addr2",
				Status: PartitionStatusInactive,
			},
			b: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusActive,
			},
			expected: false,
		},
		{
			name: "addr-equal",
			a: PartitionState{
				Term: 11,
				Addr: "addr1",
			},
			b: PartitionState{
				Term: 11,
				Addr: "addr1",
			},
			expected: false,
		},
		{
			name: "a-status-less",
			a: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusInactive,
			},
			b: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusActive,
			},
			expected: true,
		},
		{
			name: "a-status-greater",
			a: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusStopped,
			},
			b: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusInactive,
			},
			expected: false,
		},
		{
			name: "status-equal",
			a: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusStopped,
			},
			b: PartitionState{
				Term:   11,
				Addr:   "addr1",
				Status: PartitionStatusStopped,
			},
			expected: false,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := partitionStateLess(e.a, e.b)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestCombineNodeState(t *testing.T) {
	table := []struct {
		name     string
		a        NodeState
		b        NodeState
		expected NodeState
	}{
		{
			name: "both-empty",
		},
		{
			name: "a-not-empty",
			a: NodeState{
				Term:    10,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
			expected: NodeState{
				Term:    10,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
		},
		{
			name: "b-not-empty",
			b: NodeState{
				Term:    10,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
			expected: NodeState{
				Term:    10,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
		},
		{
			name: "a < b",
			a: NodeState{
				Term:    10,
				Unix:    100,
				Version: 20,
				Status:  NodeStatusStopped,
			},
			b: NodeState{
				Term:    10,
				Unix:    120,
				Version: 20,
				Status:  NodeStatusStopped,
			},
			expected: NodeState{
				Term:    10,
				Unix:    120,
				Version: 20,
				Status:  NodeStatusStopped,
			},
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := combineNodeState(e.a, e.b)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestCombineNodeStates(t *testing.T) {
	table := []struct {
		name     string
		a        map[string]NodeState
		b        map[string]NodeState
		expected map[string]NodeState
	}{
		{
			name:     "both-empty",
			expected: map[string]NodeState{},
		},
		{
			name: "a-node-non-empty",
			a: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    100,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
			expected: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    100,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
		},
		{
			name: "b-node-non-empty",
			b: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    100,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
			expected: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    100,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
		},
		{
			name: "addr-both-a-bigger",
			a: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    120,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
			b: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    100,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
			expected: map[string]NodeState{
				"addr1": {
					Term:    10,
					Unix:    120,
					Version: 20,
					Status:  NodeStatusOutOfSync,
				},
			},
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := combineNodeStates(e.a, e.b)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestCombinePartitionState(t *testing.T) {
	table := []struct {
		name     string
		a        PartitionState
		b        PartitionState
		expected PartitionState
	}{
		{
			name: "both-empty",
		},
		{
			name:     "a-term-bigger",
			a:        PartitionState{Term: 12, Addr: "addr1"},
			b:        PartitionState{Term: 10, Addr: "addr2"},
			expected: PartitionState{Term: 12, Addr: "addr1"},
		},
		{
			name:     "a-term-less",
			a:        PartitionState{Term: 10, Addr: "addr2"},
			b:        PartitionState{Term: 12, Addr: "addr1"},
			expected: PartitionState{Term: 12, Addr: "addr1"},
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := combinePartitionState(e.a, e.b)
			assert.Equal(t, e.expected, result)
		})
	}
}

func TestCombineStates(t *testing.T) {
	a := State{
		Nodes: map[string]NodeState{
			"addr1": {
				Term:    10,
				Unix:    100,
				Version: 20,
			},
		},
		Partitions: []PartitionState{
			{
				Term: 5,
				Addr: "addr1",
			},
			{
				Term: 6,
				Addr: "addr3",
			},
		},
	}

	b := State{
		Nodes: map[string]NodeState{
			"addr2": {
				Term:    30,
				Unix:    200,
				Version: 5,
			},
		},
		Partitions: []PartitionState{
			{
				Term: 6,
				Addr: "addr2",
			},
			{
				Term: 3,
				Addr: "addr4",
			},
		},
	}

	expected := State{
		Nodes: map[string]NodeState{
			"addr1": {
				Term:    10,
				Unix:    100,
				Version: 20,
			},
			"addr2": {
				Term:    30,
				Unix:    200,
				Version: 5,
			},
		},
		Partitions: []PartitionState{
			{
				Term: 6,
				Addr: "addr2",
			},

			{
				Term: 6,
				Addr: "addr3",
			},
		},
	}

	assert.Equal(t, expected, combineStates(a, b))
}
