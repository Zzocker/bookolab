package blog

// NewTestLogger create new logger for testing
func NewTestLogger() Logger {
	return testLogger{}
}

type testLogger struct{}

func (t testLogger) Infof(string, ...interface{})
func (t testLogger) Debugf(string, ...interface{})
func (t testLogger) Errorf(string, ...interface{})
