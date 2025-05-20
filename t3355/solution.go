package t3355

func isZeroArray(nums []int, queries [][]int) bool {
	dp := make([]int, len(nums))
	for _, q := range queries {
		dp[q[0]]++
		if q[1] < len(nums)-1 {
			dp[q[1]+1]--
		}
	}

	diff := 0
	for i := range dp {
		diff += dp[i]
		if nums[i] > diff {
			return false
		}
	}
	return true
}
