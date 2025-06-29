package t1498

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		nums           []int
		target         int
		expect, actual int
	)

	nums = []int{3, 5, 6, 7}
	target = 9
	expect = 4
	actual = numSubseq(nums, target)
	assert.Equal(t, expect, actual)
}
