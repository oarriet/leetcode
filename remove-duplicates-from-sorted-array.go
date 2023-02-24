package main

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastNumber := nums[0]
	i := 1

	for len(nums) > i {
		if lastNumber == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			lastNumber = nums[i]
			i++
		}
	}

	return len(nums)
}

//func main() {
//
//	//mt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
//	//fmt.Println(removeDuplicates([]int{1, 2}))
//	fmt.Println(removeDuplicates([]int{}))
//}
