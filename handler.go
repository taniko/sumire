package sumire

import (
	"encoding/json"
	"io"
)

type Handler interface {
	Handle(record Record)
}

type standardHandler struct {
	level  Level
	writer io.Writer
}

// WithStandardHandler write json format recode
func WithStandardHandler(level Level, writer io.Writer) Option {
	return WithHandler(standardHandler{
		level:  level,
		writer: writer,
	})
}

func (s standardHandler) Handle(record Record) {
	if record.Severity < s.level {
		return
	}
	b, err := json.Marshal(record)
	if err != nil {
		return
	}
	_, _ = s.writer.Write(append(b, []byte("\n")...))
}
