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
