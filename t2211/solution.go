package t2211

func countCollisions(directions string) int {
	ret := len(directions)
	for idx := 0; idx < len(directions) && directions[idx] == 'L'; idx++ {
		ret--
	}
	for idx := len(directions) - 1; idx >= 0 && directions[idx] == 'R'; idx-- {
		ret--
	}
	for idx := 0; idx < len(directions); idx++ {
		if directions[idx] == 'S' {
			ret--
		}
	}
	return ret
}
