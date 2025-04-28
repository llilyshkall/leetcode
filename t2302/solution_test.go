package t2302

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		nums           []int
		k              int64
		expect, actual int64
	)

	nums = []int{2, 1, 4, 3, 5}
	k = 10
	expect = 6
	actual = countSubarrays(nums, k)
	assert.Equal(t, expect, actual)
}
