package main

import "fmt"

func removeDuplicates(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
