package t3516

func findClosest(x int, y int, z int) int {
	if abs(x-z) == abs(y-z) {
		return 0
	}
	if abs(x-z) < abs(y-z) {
		return 1
	}
	return 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
