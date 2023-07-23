package logger

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	// Fatal Has a side effect of os.Exit(1)
	Fatal(msg string, keysAndValues ...interface{})
}
