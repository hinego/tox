package tox

func WithError(fun ...func() error) error {
	for _, f := range fun {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
