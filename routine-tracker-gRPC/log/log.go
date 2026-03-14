package log

import (
	"io"
	"log"
	"sync"
)

// A logger that can log messages
type Logger struct {
	mutext sync.Mutex
	logger *log.Logger
}

// New returns a new Logger that writes to the given writer
func New(output io.Writer) *Logger {
	return &Logger{
		logger: log.New(output, "", log.Ldate|log.Ltime),
	}
}

// Logf sends a message to the log if the severity is high enough
func (l *Logger) Logf(format string, args ...any) {
	l.mutext.Lock()
	defer l.mutext.Unlock()
	l.logger.Printf(format, args...)
}
