package sumire

import (
	"context"
	"sync"
	"time"
)

var (
	_ Logger        = (*Sumire)(nil)
	_ ContextLogger = (*Sumire)(nil)
)

type Sumire struct {
	mutex   sync.Mutex
	name    string
	options options
}

type options struct {
	filters  []RecordFilter
	handlers []Handler
}

func NewLogger(name string, opts ...Option) Sumire {
	options := options{}
	for _, opt := range opts {
		opt.apply(&options)
	}
	return Sumire{
		name:    name,
		options: options,
	}
}

func (s *Sumire) Debug(message string, context interface{}) {
	s.write(DEBUG, message, context)
}

func (s *Sumire) Info(message string, context interface{}) {
	s.write(INFO, message, context)
}

func (s *Sumire) Notice(message string, context interface{}) {
	s.write(NOTICE, message, context)
}

func (s *Sumire) Warning(message string, context interface{}) {
	s.write(WARNING, message, context)
}

func (s *Sumire) Error(message string, context interface{}) {
	s.write(ERROR, message, context)
}

func (s *Sumire) Alert(message string, context interface{}) {
	s.write(ALERT, message, context)
}

func (s *Sumire) Critical(message string, context interface{}) {
	s.write(CRITICAL, message, context)
}

func (s *Sumire) Emergency(message string, context interface{}) {
	s.write(EMERGENCY, message, context)
}

func (s *Sumire) DebugContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, DEBUG, message, context)

}

func (s *Sumire) InfoContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, INFO, message, context)
}

func (s *Sumire) NoticeContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, NOTICE, message, context)
}

func (s *Sumire) WarningContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, WARNING, message, context)
}

func (s *Sumire) ErrorContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, ERROR, message, context)
}

func (s *Sumire) CriticalContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, CRITICAL, message, context)
}

func (s *Sumire) AlertContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, ALERT, message, context)
}

func (s *Sumire) EmergencyContext(ctx context.Context, message string, context interface{}) {
	s.writeContext(ctx, EMERGENCY, message, context)
}

func (s *Sumire) write(level Level, message string, c interface{}) {
	s.writeContext(context.Background(), level, message, c)
}

func (s *Sumire) writeContext(ctx context.Context, level Level, message string, context interface{}) {
	record := Record{
		Name:      s.name,
		Severity:  level,
		Timestamp: time.Now(),
		Message:   message,
		Context:   context,
		Extra:     map[string]interface{}{},
	}
	for _, f := range s.options.filters {
		f.Filter(ctx, record)
	}
	s.writeRecord(record)
}

func (s *Sumire) writeRecord(record Record) {
	for _, handler := range s.options.handlers {
		func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()
			handler.Handle(record)
		}()
	}
}
