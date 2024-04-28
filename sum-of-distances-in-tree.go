package main

import "fmt"

type TNode struct {
	Val  int
	Prev int
	Next *TNode
}

func sumArray(arr []int) (ret int) {
	for _, v := range arr {
		ret += v
	}
	return
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ways := make([][]int, n)
	for _, way := range edges {
		ways[way[0]] = append(ways[way[0]], way[1])
		ways[way[1]] = append(ways[way[1]], way[0])
	}

	distances := make([]int, n)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		queue := &TNode{Val: i, Prev: -1}
		distances[queue.Val] = 0
		queueEnd := queue
		for queue != nil {
			for _, v := range ways[queue.Val] {
				if queue.Prev != v {
					distances[v] = distances[queue.Val] + 1
					queueEnd.Next = &TNode{Val: v, Prev: queue.Val}
					queueEnd = queueEnd.Next
				}
			}
			queue = queue.Next
		}
		ret[i] = sumArray(distances)
	}
	//fmt.Println(distances)
	return ret
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	fmt.Println(sumOfDistancesInTree(2, [][]int{{0, 1}}))
	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
}
