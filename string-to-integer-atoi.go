package main

import (
	"fmt"
	"unicode"
)

func myAtoi(s string) int {
	var ret, idx, sign, overflow = 0, 0, false, 0
	n := len(s)
	for idx < n && s[idx] == ' ' {
		idx++
	}
	if idx < n && (s[idx] == '-' || s[idx] == '+') {
		if s[idx] == '-' {
			sign = true
		}
		idx++
	}
	for idx < n && unicode.IsDigit(rune(s[idx])) {
		ret = ret*10 + int(s[idx]-'0')
		if overflow == 0 {
			if ret > 0x7fffffff {
				if sign {
					overflow = -1
				} else {
					overflow = 1
				}
			}
		}
		idx++
	}
	if sign {
		ret *= -1
	}
	if overflow > 0 {
		ret = 0x7fffffff
	}
	if overflow < 0 {
		ret = -0x80000000
	}

	return ret
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("9223372036854775808"))
}
