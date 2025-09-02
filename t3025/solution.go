package t3025

func numberOfPairs(points [][]int) int {
	ret := 0

	for j, A := range points {
		for i, B := range points {
			if i == j || !(A[0] <= B[0] && A[1] >= B[1]) {
				continue
			}

			ok := true
			for k, C := range points {
				if k == i || k == j {
					continue
				}

				if A[0] <= C[0] && C[0] <= B[0] && B[1] <= C[1] && C[1] <= A[1] {
					ok = false
					break
				}
			}
			if ok {
				ret++
			}
		}
	}
	return ret
}
