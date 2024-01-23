/*
645. Set Mismatch
Easy

You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:

Input: nums = [1,2,2,4]
Output: [2,3]
Example 2:

Input: nums = [1,1]
Output: [1,2]

Constraints:

2 <= nums.length <= 104
1 <= nums[i] <= 104
*/

package main

import (
	"fmt"
	"slices"
)

func findErrorNums(nums []int) []int {
	toReturn := make([]int, 2) //number that occurs twice , the number that is missing

	numsMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {

		if !slices.Contains(nums, i+1) {
			toReturn[1] = i + 1
		}

		if _, ok := numsMap[nums[i]]; ok {
			toReturn[0] = nums[i]
		} else {
			numsMap[nums[i]] = true
		}
	}

	return toReturn
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))                   //[2 3]
	fmt.Println(findErrorNums([]int{1, 1}))                         //[1 2]
	fmt.Println(findErrorNums([]int{3, 2, 2}))                      //[2 1]
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))             //[3,1]
	fmt.Println(findErrorNums([]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9})) //[2,10]
}
