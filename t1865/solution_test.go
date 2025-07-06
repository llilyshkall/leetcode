package t1865

import (
	"testing"
)

func TestFindSumPairs(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		ops      []func(fsp *FindSumPairs) // Последовательность операций
		validate func(t *testing.T, fsp *FindSumPairs)
	}{
		{
			name:  "Basic operations",
			nums1: []int{1, 1, 2, 2, 2, 3},
			nums2: []int{1, 4, 5, 2},
			ops: []func(fsp *FindSumPairs){
				func(fsp *FindSumPairs) {
					if got := fsp.Count(7); got != 4 {
						t.Errorf("Count(7) = %d, want 4", got)
					}
				},
				func(fsp *FindSumPairs) { fsp.Add(3, 2) },
				func(fsp *FindSumPairs) {
					if got := fsp.Count(7); got != 5 {
						t.Errorf("Count(7) after Add = %d, want 5", got)
					}
					if fsp.nums2[3] != 4 {
						t.Errorf("nums2[3] = %d, want 4", fsp.nums2[3])
					}
				},
			},
		},
		{
			name:  "Negative values",
			nums1: []int{-1, 0, 1},
			nums2: []int{-2, 3, 5},
			ops: []func(fsp *FindSumPairs){
				func(fsp *FindSumPairs) {
					if got := fsp.Count(-3); got != 1 {
						t.Errorf("Count(-3) = %d, want 1", got)
					}
				},
				func(fsp *FindSumPairs) { fsp.Add(1, -2) },
				func(fsp *FindSumPairs) {
					if got := fsp.Count(0); got != 1 {
						t.Errorf("Count(0) = %d, want 1", got)
					}
					if fsp.nums2[1] != 1 {
						t.Errorf("nums2[1] = %d, want 1", fsp.nums2[1])
					}
				},
			},
		},
		{
			name:  "Empty arrays",
			nums1: []int{},
			nums2: []int{},
			ops: []func(fsp *FindSumPairs){
				func(fsp *FindSumPairs) {
					if got := fsp.Count(10); got != 0 {
						t.Errorf("Count(10) = %d, want 0", got)
					}
				},
				func(fsp *FindSumPairs) {
					if got := fsp.Count(10); got != 0 {
						t.Errorf("Count(10) after Add = %d, want 0", got)
					}
				},
			},
		},
		{
			name:  "Multiple updates",
			nums1: []int{10},
			nums2: []int{20, 30},
			ops: []func(fsp *FindSumPairs){
				func(fsp *FindSumPairs) { fsp.Add(0, 5) },
				func(fsp *FindSumPairs) { fsp.Add(1, -10) },
				func(fsp *FindSumPairs) {
					if got := fsp.Count(30); got != 1 {
						t.Errorf("Count(30) = %d, want 1", got)
					}
					if fsp.nums2[0] != 25 || fsp.nums2[1] != 20 {
						t.Errorf("nums2 = %v, want [25, 20]", fsp.nums2)
					}
				},
			},
		},
		{
			name:  "No pairs found",
			nums1: []int{2, 4, 6},
			nums2: []int{1, 3, 5},
			ops: []func(fsp *FindSumPairs){
				func(fsp *FindSumPairs) {
					if got := fsp.Count(10); got != 0 {
						t.Errorf("Count(10) = %d, want 0", got)
					}
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsp := Constructor(tt.nums1, tt.nums2)
			for _, op := range tt.ops {
				op(&fsp)
			}
		})
	}
}
