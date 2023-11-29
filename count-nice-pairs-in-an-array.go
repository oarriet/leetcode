package main

/*1814. Count Nice Pairs in an Array

You are given an array nums that consists of non-negative integers. Let us define rev(x) as the reverse of the non-negative integer x. For example, rev(123) = 321, and rev(120) = 21. A pair of indices (i, j) is nice if it satisfies all of the following conditions:

0 <= i < j < nums.length
nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
Return the number of nice pairs of indices. Since that number can be too large, return it modulo 109 + 7.

Example 1:

Input: nums = [42,11,1,97]
Output: 2
Explanation: The two pairs are:
 - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
 - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.
Example 2:

Input: nums = [13,10,35,24,76]
Output: 4


Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 109
*/

//func countNicePairs(nums []int) int {
//	var count int
//	revNumsMap := make(map[int]int)
//
//	//let's do a pass to reverse all.
//	revNums := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		revNum, ok := revNumsMap[nums[i]]
//		if ok {
//			revNums[i] = revNum
//		} else {
//			revNums[i] = rev(nums[i])
//		}
//	}
//
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+revNums[j] == nums[j]+revNums[i] {
//				count++
//			}
//		}
//	}
//
//	return count % 1000000007
//}

func countNicePairs(nums []int) int {
	m := make(map[int]int)
	var res int
	mod := int(1e9 + 7)
	// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
	// => nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
	// So keep track of the count of nums[idx] - rev(nums[idx])
	for i := 0; i < len(nums); i++ {
		curr := nums[i] - rev(nums[i])
		res = (res + m[curr]) % mod
		m[curr]++
	}
	return res
}

func rev(num int) int {
	var res int
	for n := num; n != 0; n /= 10 {
		res = res*10 + n%10
	}
	return res
}

//func main() {
//	fmt.Println("result: ", countNicePairs([]int{42, 11, 1, 97}))
//	//fmt.Println("result: ", countNicePairs([]int{13, 10, 35, 24, 76}))
//	//fmt.Println("result: ", countNicePairs([]int{1, 1, 1, 01, 01, 01}))
//}
