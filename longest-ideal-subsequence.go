package main

import "fmt"

func longestIdealString(s string, k int) int {
	l := len(s)
	letters := make([]int, 26)
	for idx := 0; idx < l; idx++ {
		ans := s[idx] - 'a'
		left := max(int(s[idx]-'a')-k, 0)
		right := min(int(s[idx]-'a')+k, 25)
		for i := left; i <= right; i++ {
			if letters[i] > letters[ans] {
				ans = uint8(i)
			}
		}
		letters[s[idx]-'a'] = letters[ans] + 1
	}

	ret := 1
	for _, value := range letters {
		ret = max(value, ret)
	}
	return ret
}

func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
	fmt.Println(longestIdealString("pvjcci", 4))
	fmt.Println(longestIdealString("lkpkxcigcs", 6))
	fmt.Println(longestIdealString("azaza", 25))
}
