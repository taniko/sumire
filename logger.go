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
	StandardLogger interface {
		Debug(message string, context any)
		Info(message string, context any)
		Notice(message string, context any)
		Warning(message string, context any)
		Error(message string, context any)
		Critical(message string, context any)
		Alert(message string, context any)
		Emergency(message string, context any)
	}
	ContextLogger interface {
		DebugContext(ctx context.Context, message string, context any)
		InfoContext(ctx context.Context, message string, context any)
		NoticeContext(ctx context.Context, message string, context any)
		WarningContext(ctx context.Context, message string, context any)
		ErrorContext(ctx context.Context, message string, context any)
		CriticalContext(ctx context.Context, message string, context any)
		AlertContext(ctx context.Context, message string, context any)
		EmergencyContext(ctx context.Context, message string, context any)
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
