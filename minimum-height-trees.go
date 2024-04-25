package main

import "fmt"

type QNode struct {
	Val  int
	Prev int
	Next *QNode
}

func findLargestNumber(arr [][2]int) int {
	argMax := 0

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx][0] > arr[argMax][0] {
			argMax = idx
		}
	}

	return argMax
}

func bfs(ways [][]int, distances [][2]int, src int) {
	queue := &QNode{Val: src, Prev: -1}
	queueEnd := queue
	distances[src][0], distances[src][1] = 0, -1
	//fmt.Println(src, ways)
	for queue != nil {
		for _, v := range ways[queue.Val] {
			if queue.Prev != v {
				//fmt.Println("-", v, queue.Prev)
				distances[v][0] = distances[queue.Val][0] + 1
				distances[v][1] = queue.Val
				queueEnd.Next = &QNode{Val: v, Prev: queue.Val}
				queueEnd = queueEnd.Next
			}
		}
		queue = queue.Next
	}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}
	ways := make([][]int, n)
	distances := make([][2]int, n)
	for _, way := range edges {
		ways[way[0]] = append(ways[way[0]], way[1])
		ways[way[1]] = append(ways[way[1]], way[0])
	}
	bfs(ways, distances, 0)
	v1 := findLargestNumber(distances)
	//fmt.Println(distances, v1)
	//fmt.Println("---")
	bfs(ways, distances, v1)
	v2 := findLargestNumber(distances)

	maxLen := distances[v2][0]
	//fmt.Println(maxLen)
	for i := 0; i < maxLen/2; i++ {
		v2 = distances[v2][1]
	}
	if maxLen%2 == 1 {
		return []int{v2, distances[v2][1]}
	}
	//fmt.Println(distances)
	return []int{v2}
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(2, [][]int{{0, 1}}))
	fmt.Println(findMinHeightTrees(3, [][]int{{0, 1}, {0, 2}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}))
}
