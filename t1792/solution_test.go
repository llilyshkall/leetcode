package t1792

import "testing"

func Test(t *testing.T) {
	tt := []struct {
		name          string
		classes       [][]int
		extraStudents int
		expected      float64
	}{
		{
			name:          "1 class",
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			expected:      0.78333,
		},
		{
			name:          "2 class",
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			expected:      0.53485,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxAverageRatio(tc.classes, tc.extraStudents)
			if actual < tc.expected-0.00001 || actual > tc.expected+0.00001 {
				t.Errorf("expected %f, got %f", tc.expected, actual)
			}
		})
	}
}
