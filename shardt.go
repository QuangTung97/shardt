package shardt

// NodeStatus ...
type NodeStatus uint8

const (
	// NodeStatusActive ...
	NodeStatusActive NodeStatus = 0
	// NodeStatusOutOfSync ...
	NodeStatusOutOfSync NodeStatus = 1
	// NodeStatusStopped ...
	NodeStatusStopped NodeStatus = 2
)

// PartitionStatus ...
type PartitionStatus uint8

const (
	// PartitionStatusInactive ...
	PartitionStatusInactive PartitionStatus = 0
	// PartitionStatusActive ...
	PartitionStatusActive PartitionStatus = 1
	// PartitionStatusStopped ...
	PartitionStatusStopped PartitionStatus = 2
)

// NodeState ...
type NodeState struct {
	Term    uint64
	Unix    uint64
	Version uint64
	Status  NodeStatus
}

// PartitionState ...
type PartitionState struct {
	Term   uint64
	Addr   string
	Status PartitionStatus
}

// State ...
type State struct {
	Nodes      map[string]NodeState
	Partitions []PartitionState
}
