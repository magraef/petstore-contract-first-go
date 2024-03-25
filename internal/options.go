package internal

// Options represents configuration options for filter operations.
type Options struct {
	// From represents the starting id for pagination
	From int
	// Limit represents the maximum number of items to retrieve
	Limit int
}

type OptionsConfig func(options *Options)

// NewOptions creates a new Options instance with the provided configuration options.
func NewOptions(opts ...OptionsConfig) Options {
	options := Options{
		From:  0,
		Limit: 10,
	}

	for _, opts := range opts {
		opts(&options)
	}

	return options
}

func WithLimit(value *int) OptionsConfig {
	return func(options *Options) {
		if value != nil {
			options.Limit = *value
		}
	}
}

func WithFrom(value *int) OptionsConfig {
	return func(options *Options) {
		if value != nil {
			options.From = *value
		}
	}
}
