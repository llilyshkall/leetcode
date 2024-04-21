package main

import "fmt"

func dfs(start int, visited []bool, g [][]int) {
	visited[start] = true
	for _, v := range g[start] {
		if !visited[v] {
			dfs(v, visited, g)
		}
	}
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := make([][]int, n)
	for _, pair := range edges {
		g[pair[0]] = append(g[pair[0]], pair[1])
		g[pair[1]] = append(g[pair[1]], pair[0])
	}
	visited := make([]bool, n)
	dfs(source, visited, g)
	return visited[destination]
}

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	fmt.Println(validPath(10, [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}, 5, 9))
}
