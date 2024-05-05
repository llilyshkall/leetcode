package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ret := 0
	n := len(people) - 1
	for i := 0; i <= n; {
		if people[i]+people[n] <= limit {
			i++
		}
		n--
		ret++
	}
	return ret
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}
