package dialogue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertLineBreaks(t *testing.T) {

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{
			input:    "Two colourful little mushrooms. Delicious? Or deadly?",
			expected: "Two colourful little mushrooms. Delicious? Or\ndeadly?",
		},
		{
			input:    "Hello!",
			expected: "Hello!",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "This is a short sentence. No breaks.",
			expected: "This is a short sentence. No breaks.",
		},
		{
			input:    "Welcome to Sunny Valley! If you're a newcomer, that is. Otherwise, good day I suppose...",
			expected: "Welcome to Sunny Valley! If you're a newcomer,\nthat is. Otherwise, good day I suppose...",
		},
	}

	for _, tc := range tests {
		got := insertLineBreaks(tc.input)
		assert.Equal(t, tc.expected, got)
	}
}
