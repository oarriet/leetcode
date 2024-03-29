/*
1207. Unique Number of Occurrences
Easy
Topics
Companies
Hint
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.



Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true


Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/

package main

import (
	"fmt"
	"slices"
)

func uniqueOccurrences(arr []int) bool {
	occurrencesMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		occurrencesMap[arr[i]] = occurrencesMap[arr[i]] + 1
	}

	var unique []int //using it like a map
	for _, numOccur := range occurrencesMap {
		if slices.Contains(unique, numOccur) {
			return false
		} else {
			unique = append(unique, numOccur)
		}
	}

	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
