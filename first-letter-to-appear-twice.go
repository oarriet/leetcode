package main

//Given a string s consisting of lowercase English letters, return the first letter to appear twice.
//
//Note:
//
//A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
//s will contain at least one letter that appears twice.
//
//
//Example 1:
//
//Input: s = "abccbaacz"
//Output: "c"
//Explanation:
//The letter 'a' appears on the indexes 0, 5 and 6.
//The letter 'b' appears on the indexes 1 and 4.
//The letter 'c' appears on the indexes 2, 3 and 7.
//The letter 'z' appears on the index 8.
//The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.
//Example 2:
//
//Input: s = "abcdd"
//Output: "d"
//Explanation:
//The only letter that appears twice is 'd' so we return 'd'.
//
//
//Constraints:
//
//2 <= s.length <= 100
//s consists of lowercase English letters.
//s has at least one repeated letter.

//func main() {
//	fmt.Printf("abccbaacz: %c\n", repeatedCharacter("abccbaacz"))
//	fmt.Printf("abcdd: %c\n", repeatedCharacter("abcdd"))
//}

func repeatedCharacter(s string) byte {
	indexSum := 0
	var charToReturn byte

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if indexSum == 0 || j < indexSum {
					indexSum = j
					charToReturn = s[i]
				}
			}
		}
	}

	return charToReturn
}
