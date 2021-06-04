package shardt

func nodeStateLess(a NodeState, b NodeState) bool {
	if a.Term < b.Term {
		return true
	}
	if a.Term > b.Term {
		return false
	}

	if a.Unix < b.Unix {
		return true
	}
	if a.Unix > b.Unix {
		return false
	}

	if a.Version < b.Version {
		return true
	}
	if a.Version > b.Version {
		return false
	}

	return a.Status < b.Status
}

func partitionStateLess(a, b PartitionState) bool {
	if a.Term < b.Term {
		return true
	}
	if a.Term > b.Term {
		return false
	}

	if a.Addr < b.Addr {
		return true
	}
	if a.Addr > b.Addr {
		return false
	}

	return a.Status < b.Status
}

func combineNodeState(a, b NodeState) NodeState {
	if nodeStateLess(a, b) {
		return b
	}
	return a
}

func combineNodeStates(a, b map[string]NodeState) map[string]NodeState {
	result := map[string]NodeState{}
	for addr, state := range a {
		result[addr] = state
	}
	for addr, state := range b {
		previous, existed := result[addr]
		if !existed {
			result[addr] = state
		} else {
			result[addr] = combineNodeState(previous, state)
		}
	}
	return result
}

func combinePartitionState(a, b PartitionState) PartitionState {
	if partitionStateLess(a, b) {
		return b
	}
	return a
}

func combineStates(a, b State) State {
	partitions := make([]PartitionState, len(a.Partitions))
	for i, p := range a.Partitions {
		partitions[i] = combinePartitionState(p, b.Partitions[i])
	}
	return State{
		Nodes:      combineNodeStates(a.Nodes, b.Nodes),
		Partitions: partitions,
	}
}
