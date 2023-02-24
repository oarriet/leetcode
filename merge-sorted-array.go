package main

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing
the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums2Moved int
	for i := 0; i < m+n; i++ {
		for j := nums2Moved; j < n; j++ {
			if nums1[i] > nums2[j] || i >= m+nums2Moved {
				shiftRight(nums1, i, nums2[j])
				nums2Moved++
			}
		}
	}
}

func shiftRight(nums []int, positionZeroBased int, value int) {
	if len(nums) == 1 {
		nums[0] = value
		return
	}

	i := len(nums) - 1
	for i > positionZeroBased {
		nums[i] = nums[i-1]
		i--
	}
	nums[i] = value
}

//func main() {
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//
//	//test shiftRight
//	//shiftRight(nums1, 0, 9)
//	//fmt.Printf("Shifted:%v\n", nums1)
//	//
//	//shiftRight(nums1, 5, 8)
//	//fmt.Printf("Shifted:%v\n", nums1)
//
//	merge(nums1, 3, nums2, 3)
//	fmt.Printf("Result:%v\n", nums1)
//}
