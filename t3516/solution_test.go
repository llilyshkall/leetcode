package t3516

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name     string
		x, y, z  int
		expected int
	}{
		{
			name:     "test 1",
			x:        2,
			y:        7,
			z:        4,
			expected: 1,
		},
		{
			name:     "test 2",
			x:        2,
			y:        5,
			z:        6,
			expected: 2,
		},
		{
			name:     "test 3",
			x:        1,
			y:        5,
			z:        3,
			expected: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := findClosest(tc.x, tc.y, tc.z)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
