package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Key  int
	Next *Node
}

func nextStep(key int, nextKeys []int) {
	for i, d := range []int{1, 10, 100, 1000} {
		nextKeys[i*2] = key - key/d%10*d + (key/d%10+1)%10*d
		nextKeys[i*2+1] = key - key/d%10*d + (key/d%10+9)%10*d
	}
}

func openLock(deadends []string, target string) int {
	var ways [10000]int
	for _, str := range deadends {
		key, _ := strconv.Atoi(str)
		ways[key] = -1
	}
	if ways[0] == -1 {
		return -1
	}
	way := &Node{}
	end := way
	nextKeys := make([]int, 8)
	targetKey, _ := strconv.Atoi(target)
	for way != nil && way.Key != targetKey {
		nextStep(way.Key, nextKeys)
		l := ways[way.Key]
		//fmt.Println(way.Key)
		for _, key := range nextKeys {
			if ways[key] == 0 {
				ways[key] = l + 1
				end.Next = &Node{Key: key}
				end = end.Next
			}
		}
		way = way.Next
	}
	if way != nil {
		return ways[way.Key]
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	//fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}
