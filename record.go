package sumire

import (
	"context"
	"time"
)

type Record struct {
	Name      string                 `json:"name"`
	Severity  Level                  `json:"severity"`
	Timestamp time.Time              `json:"timestamp"`
	Message   string                 `json:"message"`
	Context   interface{}            `json:"context,omitempty"`
	Extra     map[string]interface{} `json:"extra,omitempty"`
}

type RecordFilter interface {
	Filter(ctx context.Context, record Record) Record
}
