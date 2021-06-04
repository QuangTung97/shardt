package shardt

type syncRemote struct {
	addr  string
	state State
}

type coreOutput struct {
	sync     []syncRemote
	expected []bool
}

type coreProcessor struct {
	partitionCount int
	options        coreOptions

	selfUnix uint64
	selfAddr string

	state State
}

func newCoreProcessor(partitions int, unix uint64, addr string, options coreOptions) *coreProcessor {
	return &coreProcessor{
		partitionCount: partitions,
		options:        options,
		selfUnix:       unix,
		selfAddr:       addr,
	}
}

func (p *coreProcessor) init() coreOutput {
	p.state = State{
		Nodes: map[string]NodeState{
			p.selfAddr: {
				Term:    1,
				Unix:    p.selfUnix,
				Version: 1,
				Status:  NodeStatusActive,
			},
		},
		Partitions: make([]PartitionState, p.partitionCount),
	}

	var sync []syncRemote
	for _, addr := range p.options.remoteAddresses {
		sync = append(sync, syncRemote{
			addr:  addr,
			state: p.state,
		})
	}

	return coreOutput{
		sync: sync,
	}
}

func (p *coreProcessor) finishSync(ok bool, state State) coreOutput {
	expected := make([]bool, p.partitionCount)
	return coreOutput{
		expected: expected,
	}
}
