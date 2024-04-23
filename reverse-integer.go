package main

import "fmt"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if ret < -0x80000000 || ret > 0x7fffffff {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
