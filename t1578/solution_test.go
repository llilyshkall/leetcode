package t1578

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name       string
		colors     string
		neededTime []int
		expected   int
	}{
		{
			name:       "test 1",
			colors:     "abaac",
			neededTime: []int{1, 2, 3, 4, 5},
			expected:   3,
		},
		{
			name:       "test 2",
			colors:     "abc",
			neededTime: []int{1, 2, 3},
			expected:   0,
		},
		{
			name:       "test 3",
			colors:     "aabaa",
			neededTime: []int{1, 2, 3, 4, 1},
			expected:   2,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := minCost(tc.colors, tc.neededTime)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
