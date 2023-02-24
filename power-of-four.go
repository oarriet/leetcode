package main

/*
Given an integer n, return true if it is a power of four. Otherwise, return false.

An integer n is a power of four, if there exists an integer x such that n == 4x.

Example 1:

Input: n = 16
Output: true
Example 2:

Input: n = 5
Output: false
Example 3:

Input: n = 1
Output: true


Constraints:

-231 <= n <= 231 - 1
*/

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%4 == 0 {
		return isPowerOfFour(n / 4)
	} else {
		return false
	}
}

//func main() {
//	println(isPowerOfFour(16))
//	println(isPowerOfFour(5))
//	println(isPowerOfFour(1))
//	println(isPowerOfFour(-2147483648))
//}
