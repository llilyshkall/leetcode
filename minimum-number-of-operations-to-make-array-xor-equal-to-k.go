package main

import "fmt"

func minOperations(nums []int, k int) int {
	xor := k
	ret := 0
	for _, val := range nums {
		xor ^= val
	}
	for i := 0; i < 32; i++ {
		ret += (xor >> i) & 1
	}
	return ret
}

func main() {
	fmt.Println(minOperations([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperations([]int{2, 0, 2, 0}, 0))
}
