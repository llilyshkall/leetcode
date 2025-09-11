package t2785

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name     string
		s        string
		expected string
	}{
		{"test1", "lEetcOde", "lEOtcede"},
		{"test2", "lYmpH", "lYmpH"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortVowels(tc.s)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
