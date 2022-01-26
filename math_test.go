package sumire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{
			a:      1,
			b:      2,
			result: 3,
		},
		{
			a:      2,
			b:      4,
			result: 6,
		},
	}

	for _, tt := range tests {
		t.Run("sum", func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			assert.Equal(t, tt.result, result)
		})
	}
}
