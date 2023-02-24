package main

/*
Given an m x n 2D binary grid `grid` which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
---
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
---

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				callDFS(grid, i, j)
			}
		}
	}

	return count
}

// callDFS from here: https://www.youtube.com/watch?v=U6-X_QOwPcs
func callDFS(grid [][]byte, i int, j int) {
	//boundary checks
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[i]) {
		return
	}

	//print island
	//printGrid(grid)

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	//call recursion
	callDFS(grid, i-1, j)
	callDFS(grid, i, j-1)
	callDFS(grid, i+1, j)
	callDFS(grid, i, j+1)

}

//func main() {
//
//	//example 1
//	//gridRow1 := []byte{'1', '1', '1', '1', '0'}
//	//gridRow2 := []byte{'1', '1', '0', '1', '0'}
//	//gridRow3 := []byte{'1', '1', '0', '0', '0'}
//	//gridRow4 := []byte{'0', '0', '0', '0', '0'}
//
//	//example 2
//	gridRow1 := []byte{'1', '1', '0', '0', '0'}
//	gridRow2 := []byte{'1', '1', '0', '0', '0'}
//	gridRow3 := []byte{'0', '0', '1', '0', '0'}
//	gridRow4 := []byte{'0', '0', '0', '1', '1'}
//
//	grid := [][]byte{gridRow1, gridRow2, gridRow3, gridRow4}
//
//	//PrintMatrixOfBytes(grid)
//	fmt.Printf("numIslands: %d", numIslands(grid))
//}
