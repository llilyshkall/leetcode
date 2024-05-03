package main

import (
	"fmt"
)

func compareVersion(version1 string, version2 string) int {
	var idx1, idx2 int
	len1, len2 := len(version1), len(version2)
	for idx1 < len1 || idx2 < len2 {
		var num1, num2 int
		if idx1 < len1 {
			for ; idx1 < len1 && version1[idx1] != '.'; idx1++ {
				num1 += num1*10 + int(version1[idx1]-'0')
			}
		}
		if idx2 < len2 {
			for ; idx2 < len2 && version2[idx2] != '.'; idx2++ {
				num2 += num2*10 + int(version2[idx2]-'0')
			}
		}
		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
		idx1++
		idx2++
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}
