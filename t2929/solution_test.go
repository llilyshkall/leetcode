package t2929

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		n, limit       int
		expect, actual int64
	)

	n = 5
	limit = 2
	expect = 3
	actual = distributeCandies(n, limit)
	assert.Equal(t, expect, actual)

	n = 3
	limit = 3
	expect = 10
	actual = distributeCandies(n, limit)
	assert.Equal(t, expect, actual)
}
