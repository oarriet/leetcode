package main

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/

func mySqrt(x int) int {

	var i int
	for i = 0; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}

	return i - 1
}

//func main() {
//	//fmt.Println(mySqrt(0))
//	fmt.Println(mySqrt(1))
//	//fmt.Println(mySqrt(4))
//	//fmt.Println(mySqrt(8))
//}
