/*
387. First Unique Character in a String
Easy
Topics
Companies
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/

package main

import (
	"fmt"
)

//brute force
//func firstUniqChar(s string) int {
//	for i := 0; i < len(s); i++ {
//		if strings.Count(s, string(s[i])) == 1 {
//			return i
//		}
//	}
//	return -1
//}

// map
func firstUniqChar(s string) int {
	m := map[byte]int{}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
