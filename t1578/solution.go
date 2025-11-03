package t1578

func minCost(colors string, neededTime []int) int {
	ret := 0
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] != colors[i+1] {
			continue
		}
		maxTime := neededTime[i]
		ret += neededTime[i]
		i++
		for i < len(colors) && colors[i-1] == colors[i] {
			maxTime = max(neededTime[i], maxTime)
			ret += neededTime[i]
			i++
		}
		ret -= maxTime
	}
	return ret
}
