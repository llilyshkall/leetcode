package t2131

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		words          []string
		expect, actual int
	)

	words = []string{"lc", "cl", "gg"}
	expect = 6
	actual = longestPalindrome(words)
	assert.Equal(t, expect, actual)

	words = []string{"ab", "ty", "yt", "lc", "cl", "ab"}
	expect = 8
	actual = longestPalindrome(words)
	assert.Equal(t, expect, actual)

	words = []string{"cc", "ll", "xx"}
	expect = 2
	actual = longestPalindrome(words)
	assert.Equal(t, expect, actual)

	words = []string{"aa", "aa", "aa"}
	expect = 6
	actual = longestPalindrome(words)
	assert.Equal(t, expect, actual)

}
