package main

import (
	"fmt"
)

func check(idx *int, s *string) bool {
	c := (*s)[*idx]
	if c == ')' || c == '}' || c == ']' {
		return true
	}
	for c == '(' || c == '{' || c == '[' {
		*idx++
		if !(check(idx, s) &&
			((c == '(' && (*s)[*idx] == ')') ||
				(c == '[' && (*s)[*idx] == ']') ||
				(c == '{' && (*s)[*idx] == '}'))) {
			return false
		}
		c = (*s)[*idx]
	}
	return true
}

func isValid(s string) bool {
	var idx int
	check(&idx, &s)
	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
