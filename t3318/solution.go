package t3318

import "sort"

func findXSum(nums []int, k int, x int) []int {
	ret := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		arr := nums[i : i+k]
		cnt := make(map[int]int)
		for _, v := range arr {
			cnt[v]++
		}

		freq := make([][2]int, 0, len(cnt))
		for key, v := range cnt {
			freq = append(freq, [2]int{key, v})
		}
		sort.Slice(freq, func(i, j int) bool {
			return freq[i][1] > freq[j][1] || freq[i][1] == freq[j][1] && freq[i][0] > freq[j][0]
		})

		xsum := 0
		for j := 0; j < x && j < len(freq); j++ {
			xsum += freq[j][0] * freq[j][1]
		}
		ret = append(ret, xsum)
	}
	return ret
}
