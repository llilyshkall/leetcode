package t3354

func countValidSelections(nums []int) int {
	leftSum, rightSum := 0, 0

	for _, n := range nums {
		rightSum += n
	}

	ret := 0

	for i := range nums {
		if nums[i] == 0 {
			if leftSum-rightSum <= 1 && leftSum-rightSum >= 0 {
				ret++
			}
			if rightSum-leftSum <= 1 && rightSum-leftSum >= 0 {
				ret++
			}
		} else {
			leftSum += nums[i]
			rightSum -= nums[i]
		}
	}
	return ret
}
