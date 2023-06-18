package main

import (
	"fmt"

)

//multi dim
func numIncreasingPaths(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = 1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 && grid[i][j] > grid[i-1][j] {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 && grid[i][j] > grid[i][j-1] {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[rows-1][cols-1]
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	numPaths := numIncreasingPaths(grid)
	fmt.Println("Number of increasing paths: ", numPaths)
}