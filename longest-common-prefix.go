package main

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.

*/

func longestCommonPrefix(strs []string) string {
	lessAmountChars := len(strs[0])

	for _, str := range strs {
		if len(str) <= lessAmountChars {
			lessAmountChars = len(str)
		}
	}

	prefix := ""
	for i := 0; i < lessAmountChars; i++ {
		charToCompare := strs[0][i]
		strsToCompare := strs[1:]
		isCharEqual := true
		for _, str := range strsToCompare {
			if charToCompare != str[i] {
				isCharEqual = false
				break
			}
		}
		if isCharEqual {
			prefix = prefix + string(charToCompare)
		} else {
			break
		}
	}

	return prefix
}

//func main() {
//	println("Result:" + longestCommonPrefix([]string{"dog", "racecar", "car"}))
//	println("Result:" + longestCommonPrefix([]string{"flower", "flow", "flight"}))
//	println("Result:" + longestCommonPrefix([]string{"a"}))
//}
