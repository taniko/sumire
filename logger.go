package sumire

import (
	"context"
	"encoding/json"
)

type Level int

var _ json.Marshaler = (*Level)(nil)

const (
	DEFAULT Level = iota
	DEBUG
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
	ALERT
	EMERGENCY
)

type (
	Logger interface {
		Debug(message string, context interface{})
		Info(message string, context interface{})
		Notice(message string, context interface{})
		Warning(message string, context interface{})
		Error(message string, context interface{})
		Critical(message string, context interface{})
		Alert(message string, context interface{})
		Emergency(message string, context interface{})
	}
	ContextLogger interface {
		DebugContext(ctx context.Context, message string, context interface{})
		InfoContext(ctx context.Context, message string, context interface{})
		NoticeContext(ctx context.Context, message string, context interface{})
		WarningContext(ctx context.Context, message string, context interface{})
		ErrorContext(ctx context.Context, message string, context interface{})
		CriticalContext(ctx context.Context, message string, context interface{})
		AlertContext(ctx context.Context, message string, context interface{})
		EmergencyContext(ctx context.Context, message string, context interface{})
	}
)

func (level Level) String() string {
	switch level {
	case DEFAULT:
		return "DEFAULT"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case NOTICE:
		return "NOTICE"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case ALERT:
		return "ALERT"
	case CRITICAL:
		return "CRITICAL"
	case EMERGENCY:
		return "EMERGENCY"
	default:
		return "DEFAULT"
	}
}

func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}
