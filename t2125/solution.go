package t2125

func numberOfBeams(bank []string) int {
	countPrev := 0
	ret := 0
	for _, b := range bank {
		count := 0
		for _, c := range b {
			count += int(c - '0')
		}
		if count != 0 {
			ret += countPrev * count
			countPrev = count
		}
	}
	return ret
}
