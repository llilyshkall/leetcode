package main

import "fmt"

func wonderfulSubstrings(word string) int64 {
	var ret int64
	d := make(map[uint16]int)
	d[0] = 1
	mask := uint16(0)

	for _, c := range word {
		mask ^= 1 << (c - 'a')
		ret += int64(d[mask])
		for i := 0; i < 10; i++ {
			ret += int64(d[mask^(1<<i)])
		}
		d[mask]++
	}

	return ret
}

func main() {
	fmt.Println(wonderfulSubstrings("aba"))
	fmt.Println(wonderfulSubstrings("aabb"))
	fmt.Println(wonderfulSubstrings("he"))
}
