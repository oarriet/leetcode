/*
739. Daily Temperatures
Medium
Topics
Companies
Hint
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is
the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which
this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/
package main

import "fmt"

// brute force, Time Limit Exceeded

//func dailyTemperatures(temperatures []int) []int {
//	r := make([]int, len(temperatures))
//
//	for i := 0; i < len(temperatures); i++ {
//		r[i] = greaterThan(temperatures[i+1:], temperatures[i])
//	}
//
//	return r
//}
//
//func greaterThan(temperatures []int, t int) int {
//	for i := 0; i < len(temperatures); i++ {
//		if temperatures[i] > t {
//			return i + 1
//		}
//	}
//	return 0
//}

// using stacks
func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}
		stack = append(stack, i)
	}

	return results
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
