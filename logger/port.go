package logger

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	// Fatal Has a side effect of panicking, should only be used if an error occurs during startup
	Fatal(msg string, keysAndValues ...interface{})
}
