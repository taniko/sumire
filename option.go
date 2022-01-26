package sumire

import (
	"context"
	"runtime"
)

const DefaultRuntimeSkip = 3

type (
	Option interface {
		apply(*options)
	}
	optionFunc func(*options)
)

func (f optionFunc) apply(options *options) {
	f(options)
}

type runtimeExtra struct {
	level Level
	skip  int
}

type appendixExtra struct {
	// key name for extra
	name string

	// key for context
	key interface{}
}

func WithHandler(handler Handler) Option {
	return optionFunc(func(o *options) {
		o.handlers = append(o.handlers, handler)
	})
}

// WithRuntimeExtra set runtime file and line to extra
func WithRuntimeExtra(level Level, skip int) Option {
	return optionFunc(func(o *options) {
		o.filters = append(o.filters, runtimeExtra{
			level: level,
			skip:  skip,
		})
	})
}

func (r runtimeExtra) Filter(_ context.Context, record Record) Record {
	if r.level > record.Severity {
		return record
	}
	if _, file, line, ok := runtime.Caller(r.skip); ok {
		record.Extra["file"] = file
		record.Extra["line"] = line
	}
	return record
}

// WithAppendixExtra set value from context to extra
func WithAppendixExtra(name string, key interface{}) Option {
	return optionFunc(func(o *options) {
		o.filters = append(o.filters, appendixExtra{
			name: name,
			key:  key,
		})
	})
}

func (a appendixExtra) Filter(ctx context.Context, record Record) Record {
	if value := ctx.Value(a.key); value != nil {
		record.Extra[a.name] = ctx.Value(a.key)
	}
	return record
}
