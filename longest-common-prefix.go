package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var idx, n int = 0, len(strs)
	check := true
	for check && idx < len(strs[0]) {
		for i := 0; i < n; i++ {
			if idx == len(strs[i]) || strs[0][idx] != strs[i][idx] {
				check = false
			}
		}
		if check {
			idx++
		}
	}
	return strs[0][:idx]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"fl", "fl", "fl"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
