package t2211

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name       string
		directions string
		expected   int
	}{
		{
			name:       "case 1",
			directions: "RLRSLL",
			expected:   5,
		},
		{
			name:       "case 2",
			directions: "LLRR",
			expected:   0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := countCollisions(tc.directions)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
