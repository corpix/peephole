package revip

// Load applies each `op` in order to fill the configuration in `v` and
// constructs a `*Revip` data-structure.
func Load(v Config, op ...Option) (*Revip, error) {
	var err error
	for _, f := range op {
		err = f(v)
		if err != nil {
			return nil, err
		}
	}

	return New(v), nil
}
