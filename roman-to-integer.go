// нужно оптимизировать
package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	var dict = map[rune]int{
		'M': 1000,
		'm': 900,
		'D': 500,
		'd': 400,
		'C': 100,
		'c': 90,
		'L': 50,
		'l': 40,
		'X': 10,
		'x': 9,
		'V': 5,
		'v': 4,
		'I': 1,
	}
	s = strings.ReplaceAll(s, "CM", "m")
	s = strings.ReplaceAll(s, "CD", "d")
	s = strings.ReplaceAll(s, "XC", "c")
	s = strings.ReplaceAll(s, "XL", "l")
	s = strings.ReplaceAll(s, "IX", "x")
	s = strings.ReplaceAll(s, "IV", "v")
	ret := 0
	for _, c := range []rune(s) {
		for key, num := range dict {
			if key == c {
				ret += num
				break
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
