package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractNumberSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple suffix",
			input:    "eval-3",
			expected: 3,
		},
		{
			name:     "larger number",
			input:    "myversion-123",
			expected: 123,
		},
		{
			name:     "no suffix",
			input:    "eval",
			expected: -1,
		},
		{
			name:     "underscore instead of dash",
			input:    "eval_3",
			expected: -1,
		},
		{
			name:     "number in middle",
			input:    "eval-3-test",
			expected: -1,
		},
		{
			name:     "empty string",
			input:    "",
			expected: -1,
		},
		{
			name:     "only dash and number",
			input:    "-42",
			expected: 42,
		},
		{
			name:     "multiple dashes",
			input:    "my-version-5",
			expected: 5,
		},
		{
			name:     "zero suffix",
			input:    "eval-0",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractNumberSuffix(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
