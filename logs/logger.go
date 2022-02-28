package logs

// Logger is the main interface of this package, it's the only interface
// available to interact with every logger implementation, that's to keep the
// usage easily integrated in code and tests.
type Logger interface {
	// Tag returns a new Logger with the extra tag added to it. Tags are logged
	// on every line as a prefix of the message in the form [name:value]
	Tag(name string, value interface{}) Logger

	// Errorf will log a message applying args to the format string in the same
	// way fmt.Sprintf works but applying the [ERROR] prefix to the message
	Errorf(format string, args ...interface{})

	// Warnf will log a message applying args to the format string in the same
	// way fmt.Sprintf works but applying the [WARN] prefix to the message
	Warnf(format string, args ...interface{})

	// Infof will log a message applying args to the format string in the same
	// way fmt.Sprintf works but applying the [INFO] prefix to the message
	Infof(format string, args ...interface{})
}
