package main

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0, 1
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		j := 1
		total := 0
		num := 0
		den := 0
		maxFrac := 0.0
		for i := 0; i < n; i++ {
			for j < n && arr[i] >= arr[j]*mid {
				j++
			}
			total += n - j
			if j < n && maxFrac < float64(arr[i])/float64(arr[j]) {
				maxFrac = float64(arr[i]) / float64(arr[j])
				num = i
				den = j
			}
		}
		if total == k {
			return []int{arr[num], arr[den]}
		}

		if total > k {
			right = mid
		} else {
			left = mid
		}
	}

	return []int{0, 0}
}

func main() {

}
