package internal

type Options struct {
	From  int
	Limit int
}

type OptionsConfig func(options *Options)

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
