package t1498

import "sort"

func numSubseq(nums []int, target int) int {
	const Mod = 1e9 + 7
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ret := 0

	pows2 := make([]int, len(nums))
	pows2[0] = 1
	for i := 1; i < len(pows2); i++ {
		pows2[i] = (pows2[i-1] * 2) % Mod
	}

	for left <= right {
		if nums[left]+nums[right] <= target {
			ret = (ret + pows2[right-left]) % Mod
			left++
		} else {
			right--
		}
	}
	return ret
}
