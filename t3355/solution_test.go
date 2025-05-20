package t3355

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		nums           []int
		queries        [][]int
		expect, actual bool
	)

	nums = []int{1, 0, 1}
	queries = [][]int{{0, 2}}
	expect = true
	actual = isZeroArray(nums, queries)
	assert.Equal(t, expect, actual)

	nums = []int{4, 3, 2, 1}
	queries = [][]int{{1, 3}, {0, 2}}
	expect = false
	actual = isZeroArray(nums, queries)
	assert.Equal(t, expect, actual)
}
