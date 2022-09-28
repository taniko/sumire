package sumire

import (
	"context"
	"os"
	"sync"
	"time"
)

var (
	_ StandardLogger = (*Logger)(nil)
	_ ContextLogger  = (*Logger)(nil)
)

type Logger struct {
	mutex   sync.Mutex
	name    string
	options options
}

type options struct {
	filters  []RecordFilter
	handlers []Handler
	exit     *Level
}

func NewLogger(name string, opts ...Option) *Logger {
	options := options{}
	for _, opt := range opts {
		opt.apply(&options)
	}
	return &Logger{
		name:    name,
		options: options,
	}
}
func (l *Logger) Debug(message string, context any) {
	l.write(DEBUG, message, context)
}

func (l *Logger) Info(message string, context any) {
	l.write(INFO, message, context)
}

func (l *Logger) Notice(message string, context any) {
	l.write(NOTICE, message, context)
}

func (l *Logger) Warning(message string, context any) {
	l.write(WARNING, message, context)
}

func (l *Logger) Error(message string, context any) {
	l.write(ERROR, message, context)
}

func (l *Logger) Alert(message string, context any) {
	l.write(ALERT, message, context)
}

func (l *Logger) Critical(message string, context any) {
	l.write(CRITICAL, message, context)
}

func (l *Logger) Emergency(message string, context any) {
	l.write(EMERGENCY, message, context)
}

func (l *Logger) DebugContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, DEBUG, message, context)

}

func (l *Logger) InfoContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, INFO, message, context)
}

func (l *Logger) NoticeContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, NOTICE, message, context)
}

func (l *Logger) WarningContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, WARNING, message, context)
}

func (l *Logger) ErrorContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, ERROR, message, context)
}

func (l *Logger) CriticalContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, CRITICAL, message, context)
}

func (l *Logger) AlertContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, ALERT, message, context)
}

func (l *Logger) EmergencyContext(ctx context.Context, message string, context any) {
	l.writeContext(ctx, EMERGENCY, message, context)
}

func (l *Logger) write(level Level, message string, c any) {
	l.writeContext(context.TODO(), level, message, c)
}

func (l *Logger) writeContext(ctx context.Context, level Level, message string, context any) {
	record := Record{
		Name:      l.name,
		Severity:  level,
		Timestamp: time.Now(),
		Message:   message,
		Context:   context,
		Extra:     map[string]interface{}{},
	}
	for _, f := range l.options.filters {
		f.Filter(ctx, record)
	}
	l.writeRecord(record)
	if l.options.exit != nil && level >= *l.options.exit {
		os.Exit(1)
	}
}

func (l *Logger) writeRecord(record Record) {
	for _, handler := range l.options.handlers {
		func() {
			l.mutex.Lock()
			defer l.mutex.Unlock()
			handler.Handle(record)
		}()
	}
}
