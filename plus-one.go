package main

/*

 */

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0 && carry; i-- {
		if digits[i] < 9 {
			carry = false
			digits[i] = digits[i] + 1
		} else {
			carry = true
			digits[i] = 0
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}

//func main() {
//	fmt.Println(plusOne([]int{1, 2, 3}))
//	fmt.Println(plusOne([]int{4, 3, 2, 1}))
//	fmt.Println(plusOne([]int{9}))
//	fmt.Println(plusOne([]int{9, 9}))
//
//}
