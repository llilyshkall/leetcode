package main

import "fmt"

func isMatchByte(s []byte, p []byte) bool {
	if len(p) > 1 && p[1] == '*' {
		ret := isMatchByte(s[0:], p[2:])
		for i := 1; i < len(s)+1 && (p[0] == '.' || p[0] == s[i-1]) && !ret; i++ {
			ret = isMatchByte(s[i:], p[2:])
		}
		return ret
	} else if len(p) == 0 || len(s) == 0 {
		return len(s) == len(p)
	} else if p[0] == '.' || p[0] == s[0] {
		return isMatchByte(s[1:], p[1:])
	} else {
		return false
	}
}

func isMatch(s string, p string) bool {
	return isMatchByte([]byte(s), []byte(p))
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("a", "ab*"))
}
