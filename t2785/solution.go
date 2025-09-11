package t2785

func isVowel(c byte) int {
	switch c {
	case 'a':
		return 5
	case 'e':
		return 6
	case 'i':
		return 7
	case 'o':
		return 8
	case 'u':
		return 9
	case 'A':
		return 0
	case 'E':
		return 1
	case 'I':
		return 2
	case 'O':
		return 3
	case 'U':
		return 4
	}
	return -1
}

func sortVowels(s string) string {
	ret := make([]byte, len(s))
	vowels := make([]int, 10)
	for i := 0; i < len(s); i++ {
		if idx := isVowel(s[i]); idx == -1 {
			ret[i] = s[i]
		} else {
			vowels[idx]++
		}
	}
	idx := 0
	vowelsAll := "AEIOUaeiou"
	for i := 0; i < len(s); i++ {
		if ret[i] == 0 {
			for vowels[idx] == 0 {
				idx++
			}
			ret[i] = vowelsAll[idx]
			vowels[idx]--
		}
	}
	return string(ret)
}
