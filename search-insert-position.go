package main

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func searchInsert(nums []int, target int) int {

	for i, num := range nums {
		if num == target {
			return i
		}
		if target < num {
			return i
		}
	}
	return len(nums)
}

//func main() {
//	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
//	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
//	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
//}
