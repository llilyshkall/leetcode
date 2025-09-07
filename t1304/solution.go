package t1304

func sumZero(n int) []int {
	ret := make([]int, n)
	for i := 1; i <= n/2; i++ {
		ret[(i-1)*2] = i
		ret[(i-1)*2+1] = -i
	}
	return ret
}
