package t3025

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "test 1",
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: 0,
		},
		{
			name:     "test 2",
			points:   [][]int{{6, 2}, {4, 4}, {2, 6}},
			expected: 2,
		},
		{
			name:     "test 3",
			points:   [][]int{{3, 1}, {1, 3}, {1, 1}},
			expected: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := numberOfPairs(tc.points)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
