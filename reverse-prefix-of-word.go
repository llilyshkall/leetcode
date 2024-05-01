package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	idx := 0
	n := len(word)
	for idx < n && word[idx] != ch {
		idx++
	}
	if idx == n {
		return word
	} else {
		symbols := []byte(word[0 : idx+1])

		for i := 0; i <= idx/2; i++ {
			symbols[i], symbols[idx-i] = symbols[idx-i], symbols[i]
		}
		return string(symbols) + word[idx+1:]
	}
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}
