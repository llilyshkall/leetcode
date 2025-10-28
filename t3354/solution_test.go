package t3354

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 0, 2, 0, 3},
			expect: 2,
		},
		{
			name:   "test 2",
			nums:   []int{2, 3, 4, 0, 4, 1, 0},
			expect: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := countValidSelections(tc.nums)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
