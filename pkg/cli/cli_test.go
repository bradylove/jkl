package cli_test

type nopLogger struct{}

func (nopLogger) Printf(string, ...interface{}) {
	panic("not implemented")
}

func (nopLogger) Fatalf(string, ...interface{}) {
	panic("not implemented")
}
