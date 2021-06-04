package shardt

type coreOptions struct {
	remoteAddresses []string
}

// Option ...
type Option func(opts *coreOptions)

func defaultCoreOptions() coreOptions {
	return coreOptions{}
}

func computeOptions(options ...Option) coreOptions {
	opts := defaultCoreOptions()
	for _, o := range options {
		o(&opts)
	}
	return opts
}

// AddRemoteAddress ...
func AddRemoteAddress(addr string) Option {
	return func(opts *coreOptions) {
		opts.remoteAddresses = append(opts.remoteAddresses, addr)
	}
}
