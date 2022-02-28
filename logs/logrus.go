package logs

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"runtime"
	"sort"
	"strings"

	libruntime "github.com/NaujOyamat/infinity/runtime"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	entry *logrus.Entry
}

// NewLogrusLogger creates a new Logger that uses https://github.com/sirupsen/logrus
// implementation.
func NewLogrusLogger() Logger {

	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(logrusFormatter{
		timeFormat: "2006-01-02 15:04:05.14Z07:00",
	})

	logger.SetLevel(logrus.InfoLevel)

	return LogrusLogger{
		entry: logrus.NewEntry(logger),
	}
}

// Tag returns a new Logger with the extra tag added to it. Tags are logged
// on every line as a prefix of the message in the form [name:value]
func (l LogrusLogger) Tag(name string, value interface{}) Logger {
	l.entry = l.entry.WithField(name, value)
	return l
}

// Errorf will log a message applying args to the format string in the same
// way fmt.Sprintf works but applying the [ERROR] prefix to the message
func (l LogrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Warnf will log a message applying args to the format string in the same
// way fmt.Sprintf works but applying the [WARN] prefix to the message
func (l LogrusLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Infof will log a message applying args to the format string in the same
// way fmt.Sprintf works but applying the [INFO] prefix to the message
func (l LogrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

type logrusFormatter struct {
	timeFormat string
}

func (f logrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry == nil {
		return []byte{}, errors.New("nil logrus.Entry to format")
	}

	caller, file := prettifyCaller(entry.Caller)

	var buf bytes.Buffer
	_, err := fmt.Fprintf(&buf, `[%s] [%s] [%s] [%s] `,
		entry.Time.Format(f.timeFormat),
		strings.ToUpper(entry.Level.String()),
		caller, file,
	)
	if err != nil {
		return nil, err
	}

	tags := make([]string, len(entry.Data))
	for k, v := range entry.Data {
		tags = append(tags, fmt.Sprintf(`[%s:%v] `, k, v))
	}
	sort.Strings(tags)
	for _, tag := range tags {
		_, err := fmt.Fprint(&buf, tag)
		if err != nil {
			return nil, err
		}
	}

	n, err := buf.WriteString(entry.Message)
	if err != nil {
		return nil, err
	}
	if len(entry.Message) != n {
		return nil, fmt.Errorf("can't write the complete log message, want to write %d bytes, but wrote %d bytes", len(entry.Message), n)
	}

	n, err = buf.WriteRune('\n')
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, fmt.Errorf("can't write end of line, want to write 1 byte, but wrote %d bytes", n)
	}

	return buf.Bytes(), nil
}

func prettifyCaller(frame *runtime.Frame) (string, string) {
	// we need the next frame because logrus reports the caller as being
	// always LogrusLogger on this file
	frame = libruntime.Next(frame)
	if frame == nil {
		return "caller_not_found", "file_not_found"
	}
	fn := strings.Split(fmt.Sprintf("%s()", frame.Function), "/")
	return fn[len(fn)-1], fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
}
