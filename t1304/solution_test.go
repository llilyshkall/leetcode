package t1304

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tt := []struct {
		name     string
		n        int
		expected []int
	}{
		{"n=1", 1, []int{0}},
		{"n=2", 2, []int{1, -1}},
		{"n=3", 3, []int{1, -1, 0}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := sumZero(tc.n)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}
