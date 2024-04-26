package main

import "fmt"

func twoMinNumber(n int, arr []int) (int, int) {
	if n == 1 {
		return 0, 0
	}
	argMin1 := 0
	argMin2 := 1
	if arr[argMin1] > arr[argMin2] {
		argMin1 = 1
		argMin2 = 0
	}
	for i := 2; i < n; i++ {
		if arr[i] < arr[argMin1] {
			argMin2 = argMin1
			argMin1 = i
		} else if arr[i] < arr[argMin2] {
			argMin2 = i
		}
	}
	return argMin1, argMin2
}

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	for i := 1; i < n; i++ {
		argMin1, argMin2 := twoMinNumber(n, grid[i-1])
		for j := 0; j < n; j++ {
			if argMin1 != j {
				grid[i][j] += grid[i-1][argMin1]
			} else {
				grid[i][j] += grid[i-1][argMin2]
			}
		}
	}
	argMin, _ := twoMinNumber(n, grid[n-1])
	return grid[n-1][argMin]
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{7}}))
	fmt.Println(minFallingPathSum([][]int{{-73, 61, 43, -48, -36}, {3, 30, 27, 57, 10}, {96, -76, 84, 59, -15}, {5, -49, 76, 31, -7}, {97, 91, 61, -46, 67}}))
}
