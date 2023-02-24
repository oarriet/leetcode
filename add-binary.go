package main

/*
Given two binary strings a and b, return their sum as a binary string.

0 + 0 = 0
0 + 1 = 1
1 + 0 = 1
1 + 1 =10
*/

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		b, a = a, b
	}

	b = fillWithZeros(b, len(a)-len(b))

	i := 0
	carry := false
	total := ""
	for i <= len(a)-1 {
		var result string
		carry, result = sum(a[len(a)-1-i], b[len(b)-1-i], carry)
		total = result + total
		i++
	}

	if carry {
		total = "1" + total
	}

	return total
}

func fillWithZeros(val string, zeroAmount int) string {
	if zeroAmount == 0 {
		return val
	}

	for zeroAmount > 0 {
		val = "0" + val
		zeroAmount--
	}
	return val
}

/*
0 %!s(int32=48)
1 %!s(int32=49)
*/
func sum(a uint8, b uint8, carry bool) (bool, string) {
	if carry {
		if a == 48 && b == 48 {
			return false, "1"
		}

		if a == 48 && b == 49 {
			return true, "0"
		}

		if a == 49 && b == 48 {
			return true, "0"
		}

		return true, "1"
	} else {
		if a == 48 && b == 48 {
			return false, "0"
		}

		if a == 48 && b == 49 {
			return false, "1"
		}

		if a == 49 && b == 48 {
			return false, "1"
		}

		return true, "0"
	}

}

//func main() {
//	fmt.Println(addBinary("11", "1"))      //100
//	fmt.Println(addBinary("1010", "1011")) //10101
//	fmt.Println(addBinary("1111", "1111")) //11110
//}
