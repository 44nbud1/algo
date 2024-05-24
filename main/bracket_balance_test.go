package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBracketBalance(t *testing.T) {

	testCase := []struct {
		name     string
		bracket  string
		expected string
	}{
		{
			name:     "success - 1",
			bracket:  "{ [ ( )]}",
			expected: "YES",
		},
		{
			name:     "success with comma - 2",
			bracket:  "{ [ ( ),],}",
			expected: "YES",
		},
		{
			name:     "failed",
			bracket:  "{ [ ( ] ) }",
			expected: "NO",
		},
		{
			name:     "success - 3",
			bracket:  "{ ( ( [ ] ) [ ] ) [ ] }",
			expected: "YES",
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {

			actual := BracketBalance(tc.bracket)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
