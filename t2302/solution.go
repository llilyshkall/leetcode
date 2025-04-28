package t2302

func countSubarrays(nums []int, k int64) int64 {
	var (
		sum, ret int64
	)

	for left, right := 0, 0; right < len(nums); right++ {
		sum += int64(nums[right])
		for ; left <= right && sum*int64(right-left+1) >= k; left++ {
			sum -= int64(nums[left])
		}
		ret += int64(right - left + 1)
	}
	return ret
}
