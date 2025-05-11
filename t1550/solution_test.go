package t1550

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		nums           []int
		expect, actual bool
	)

	nums = []int{2, 6, 4, 1}
	expect = false
	actual = threeConsecutiveOdds(nums)
	assert.Equal(t, expect, actual)

	nums = []int{1, 1, 1}
	expect = true
	actual = threeConsecutiveOdds(nums)
	assert.Equal(t, expect, actual)

	nums = []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	expect = true
	actual = threeConsecutiveOdds(nums)
	assert.Equal(t, expect, actual)
}
