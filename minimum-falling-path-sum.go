/*
931. Minimum Falling Path Sum
Medium
Topics
Companies
Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).



Example 1:


Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Explanation: There are two falling paths with a minimum sum as shown.
Example 2:


Input: matrix = [[-19,57],[-40,-5]]
Output: -59
Explanation: The falling path with a minimum sum is shown.


Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100
*/

package main

import (
	"fmt"
	"slices"
)

// based in https://www.youtube.com/watch?v=b_F3mz9l-uQ
func minFallingPathSum(matrix [][]int) int {
	N := len(matrix)
	for r := 1; r < N; r++ {
		for c := 0; c < N; c++ {
			value := matrix[r][c]
			switch c {
			case 0:
				matrix[r][c] = min(value+matrix[r-1][c], value+matrix[r-1][c+1])
			case N - 1:
				matrix[r][c] = min(value+matrix[r-1][c-1], value+matrix[r-1][c])
			default:
				matrix[r][c] = min(value+matrix[r-1][c-1], value+matrix[r-1][c], value+matrix[r-1][c+1])
			}
		}
	}

	return slices.Min(matrix[N-1])
}

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))
}
