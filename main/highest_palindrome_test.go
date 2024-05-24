package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestPalindrome(t *testing.T) {

	testCase := []struct {
		name        string
		palindrome  string
		replacement int
		expected    string
	}{
		{
			name:        "success - 1",
			palindrome:  "3943",
			replacement: 1,
			expected:    "3993",
		},
		{
			name:        "success - 2",
			palindrome:  "932239",
			replacement: 2,
			expected:    "992299",
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {

			actual := HighestPalindrome(tc.palindrome, tc.replacement)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
