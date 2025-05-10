package t2918

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		nums1, nums2   []int
		expect, actual int64
	)

	nums1 = []int{3, 2, 0, 1, 0}
	nums2 = []int{6, 5, 0}
	expect = int64(12)
	actual = minSum(nums1, nums2)
	assert.Equal(t, expect, actual)

	nums1 = []int{2, 0, 2, 0}
	nums2 = []int{1, 4}
	expect = int64(-1)
	actual = minSum(nums1, nums2)
	assert.Equal(t, expect, actual)
}
