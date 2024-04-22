// нужно оптимизировать
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var idx1, idx2, prev1, prev2 int
	n, m := len(nums1), len(nums2)
	sumlen := (n + m + 2) / 2
	for idx1+idx2 < sumlen {
		prev2 = prev1
		if idx1 < n && idx2 < m {
			if nums1[idx1] < nums2[idx2] {
				prev1 = nums1[idx1]
				idx1++
			} else {
				prev1 = nums2[idx2]
				idx2++
			}
		} else {
			if idx1 < n {
				prev1 = nums1[idx1]
				idx1++
			} else {
				prev1 = nums2[idx2]
				idx2++
			}
		}
		fmt.Println("-", prev1)
	}
	if (n+m)%2 == 0 {
		return float64(prev1+prev2) / 2.0
	}
	return float64(prev1)
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
