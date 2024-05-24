package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeightedStrings(t *testing.T) {

	testCase := []struct {
		name     string
		alpha    string
		queries  []int
		expected []string
	}{
		{
			name:     "success",
			alpha:    "abbcccd",
			queries:  []int{1, 3, 9, 8},
			expected: []string{"Yes", "Yes", "Yes", "No"},
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {

			actual := WeightedStrings(tc.alpha, tc.queries)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
