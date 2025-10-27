package t2125

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name   string
		bank   []string
		expect int
	}{
		{
			name:   "test 1",
			bank:   []string{"011001", "000000", "010100", "001000"},
			expect: 8,
		},
		{
			name:   "test 2",
			bank:   []string{"000", "111", "000"},
			expect: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := numberOfBeams(tc.bank)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
