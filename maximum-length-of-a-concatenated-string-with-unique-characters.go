/*
1239. Maximum Length of a Concatenated String with Unique Characters
Medium
Topics
Companies
Hint
You are given an array of strings arr. A string s is formed by the concatenation of a subsequence of arr that has unique characters.

Return the maximum possible length of s.

A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.



Example 1:

Input: arr = ["un","iq","ue"]
Output: 4
Explanation: All the valid concatenations are:
- ""
- "un"
- "iq"
- "ue"
- "uniq" ("un" + "iq")
- "ique" ("iq" + "ue")
Maximum length is 4.
Example 2:

Input: arr = ["cha","r","act","ers"]
Output: 6
Explanation: Possible longest valid concatenations are "chaers" ("cha" + "ers") and "acters" ("act" + "ers").
Example 3:

Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
Output: 26
Explanation: The only string in arr has all 26 characters.


Constraints:

1 <= arr.length <= 16
1 <= arr[i].length <= 26
arr[i] contains only lowercase English letters.
*/

package main

import (
	"fmt"
	"strings"
)

func maxLength(arr []string) int {
	return backtrack(0, "", arr)
}

func backtrack(i int, s string, arr []string) int {
	if i == len(arr) {
		return len(s)
	}
	if hasUniqueChars(s + arr[i]) {
		return max(backtrack(i+1, s+arr[i], arr), backtrack(i+1, s, arr))
	} else {
		return backtrack(i+1, s, arr)
	}
}

func hasUniqueChars(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(s[i+1:], rune(s[i])) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
}
