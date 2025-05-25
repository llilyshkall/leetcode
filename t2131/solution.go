package t2131

func longestPalindrome(words []string) int {
	m := make(map[string]int)

	for _, s := range words {
		m[s]++
	}

	ret := 0
	flag := false
	for s, k := range m {
		reverse := string([]byte{s[1], s[0]})

		if reverse == s && k%2 == 1 {
			if !flag {
				flag = true
				ret += 2
			}
			ret += (k - k%2) * 2
		} else {
			ret += min(k, m[reverse]) * 2
		}
	}
	return ret
}
