package logger

type TestLogger struct {
	Messages []string // At the end of the test you can check if the loggers you expected were included.
}

func NewTestLogger() *TestLogger {
	return &TestLogger{
		Messages: make([]string, 0),
	}
}

func (t *TestLogger) Debug(msg string, _ ...interface{}) {
	t.Messages = append(t.Messages, msg)
}

func (t *TestLogger) Info(msg string, _ ...interface{}) {
	t.Messages = append(t.Messages, msg)
}

func (t *TestLogger) Warn(msg string, _ ...interface{}) {
	t.Messages = append(t.Messages, msg)
}

func (t *TestLogger) Error(msg string, _ ...interface{}) {
	t.Messages = append(t.Messages, msg)
}

func (t *TestLogger) Fatal(msg string, _ ...interface{}) {
	t.Messages = append(t.Messages, msg)
}
