package t3318

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name     string
		nums     []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:        6,
			x:        2,
			expected: []int{6, 10, 12},
		},
		{
			name:     "test2",
			nums:     []int{3, 8, 7, 8, 7, 5},
			k:        2,
			x:        2,
			expected: []int{11, 15, 15, 15, 12},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := findXSum(tc.nums, tc.k, tc.x)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
