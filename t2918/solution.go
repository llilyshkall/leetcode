package t2918

func minSum(nums1 []int, nums2 []int) int64 {
	var (
		sum1, sum2             int64
		countZero1, countZero2 int
	)
	sum1, countZero1 = sum(nums1)
	sum2, countZero2 = sum(nums2)
	minSum1 := sum1 + int64(countZero1)
	minSum2 := sum2 + int64(countZero2)
	if minSum1 < minSum2 {
		if countZero1 > 0 {
			return minSum2
		} else {
			return -1
		}
	}
	if minSum1 > minSum2 {
		if countZero2 > 0 {
			return minSum1
		} else {
			return -1
		}
	}
	return minSum1
}

func sum(nums []int) (int64, int) {
	ret := int64(0)
	countZero := 0
	for _, num := range nums {
		ret += int64(num)
		if num == 0 {
			countZero++
		}
	}
	return ret, countZero
}
