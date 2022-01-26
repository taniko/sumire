package sumire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		level Level
		str   string
	}{
		{
			level: INFO,
			str:   "INFO",
		}, {
			level: WARNING,
			str:   "WARNING",
		},
	}

	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			assert.Equal(t, tt.str, tt.level.String())
		})
	}
}

func TestLevel_MarshalJSON(t *testing.T) {
	tests := []struct {
		level Level
		bytes []byte
	}{
		{
			level: INFO,
			bytes: []byte("\"INFO\""),
		},
		{
			level: WARNING,
			bytes: []byte("\"WARNING\""),
		},
		{
			level: 100,
			bytes: []byte("\"DEFAULT\""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			b, err := tt.level.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.bytes, b)
		})
	}
}
