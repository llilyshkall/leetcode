package t3403

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		word           string
		numFriends     int
		expect, actual string
	)

	word = "dbca"
	numFriends = 2
	expect = "dbc"
	actual = answerString(word, numFriends)
	assert.Equal(t, expect, actual)

	word = "gggg"
	numFriends = 4
	expect = "g"
	actual = answerString(word, numFriends)
	assert.Equal(t, expect, actual)
}
