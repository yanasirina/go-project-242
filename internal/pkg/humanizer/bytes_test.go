package humanizer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHumanizeBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		base     float64
		expected string
	}{
		{
			name:     "Zero bytes",
			input:    0,
			base:     1024,
			expected: "0B",
		},
		{
			name:     "Less than 10 bytes",
			input:    5,
			base:     1024,
			expected: "5B",
		},
		{
			name:     "Exactly 10 bytes",
			input:    10,
			base:     1024,
			expected: "10B",
		},
		{
			name:     "Bytes to kilobytes",
			input:    1500,
			base:     1024,
			expected: "1.5KB",
		},
		{
			name:     "Bytes to megabytes",
			input:    2_621_440,
			base:     1024,
			expected: "2.5MB",
		},
		{
			name:     "Bytes to gigabytes",
			input:    3_758_096_384,
			base:     1024,
			expected: "3.5GB",
		},
		{
			name:     "Round values",
			input:    1_048_576,
			base:     1024,
			expected: "1.0MB",
		},
		{
			name:     "Large values",
			input:    1_099_511_627_776,
			base:     1024,
			expected: "1.0TB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanizeBytes(tt.input, tt.base)
			require.Equal(t, tt.expected, result)
		})
	}
}
