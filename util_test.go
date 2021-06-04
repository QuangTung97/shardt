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
