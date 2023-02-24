package main

//func main() {
//	fmt.Printf("Results: %v\n", loop(0, 2, 10))
//	fmt.Printf("Results: %v\n", loop(5, 3, 5))
//}

func loop(a, b, n int) []int {
	results := make([]int, n)

	for i := 0; i < n; i++ {

		var result int
		if i == 0 {
			result = a + (expFunc(2, i) * b)
		} else {
			result = results[i-1] + (expFunc(2, i) * b)
		}

		results[i] = result
	}

	return results
}

func expFunc(num, exp int) int {
	if exp == 0 {
		return 1
	}
	result := num

	for i := 0; i < exp-1; i++ {
		result = result * num
	}

	return result
}
