package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
1481. Least Number of Unique Integers after K Removals
Medium
Topics
Companies
Hint
Given an array of integers arr and an integer k. Find the least number of unique integers after removing exactly k elements.



Example 1:

Input: arr = [5,5,4], k = 1
Output: 1
Explanation: Remove the single 4, only 5 is left.
Example 2:
Input: arr = [4,3,1,1,3,3,2], k = 3
Output: 2
Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.


Constraints:

1 <= arr.length <= 10^5
1 <= arr[i] <= 10^9
0 <= k <= arr.length
*/

func findLeastNumOfUniqueInts(arr []int, k int) int {
	mapOccurrence := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		mapOccurrence[arr[i]]++
	}

	var listOccurrence []int
	for mKey, _ := range mapOccurrence {
		listOccurrence = append(listOccurrence, mKey)
	}

	slices.SortFunc(listOccurrence, func(a, b int) int {
		return cmp.Compare(mapOccurrence[a], mapOccurrence[b])
	})

	for k > 0 {
		if mapOccurrence[listOccurrence[0]] > k {
			k = k - mapOccurrence[listOccurrence[0]]
		} else {
			k = k - mapOccurrence[listOccurrence[0]]
			listOccurrence = listOccurrence[1:]
		}
	}

	return len(listOccurrence)
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))             //1
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3)) //2
}
