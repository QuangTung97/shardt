package shardt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllocatePartitions(t *testing.T) {
	table := []struct {
		name     string
		count    int
		addrs    []string
		previous []string
		expected []string
	}{
		{
			name:     "single-addr",
			count:    3,
			addrs:    []string{"addr1"},
			previous: []string{"", "", ""},
			expected: []string{"addr1", "addr1", "addr1"},
		},
		{
			name:     "two-addrs",
			count:    3,
			addrs:    []string{"addr1", "addr2"},
			previous: []string{"", "", ""},
			expected: []string{"addr1", "addr1", "addr2"},
		},
		{
			name:     "previous-only-addr2",
			count:    3,
			addrs:    []string{"addr1", "addr2"},
			previous: []string{"addr2", "addr2", "addr2"},
			expected: []string{"addr2", "addr1", "addr1"},
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result := allocatePartitions(e.count, e.addrs, e.previous)
			assert.Equal(t, e.expected, result)
		})
	}
}
