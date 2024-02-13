package main

import "fmt"

/*
2108. Find First Palindromic String in the Array
Easy
Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.

Example 1:

Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
Example 2:

Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar".
Example 3:

Input: words = ["def","ghi"]
Output: ""
Explanation: There are no palindromic strings, so the empty string is returned.


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists only of lowercase English letters.
*/

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}
	return ""
}

func isPalindrome(w string) bool {
	l := len(w)
	for i := 0; i < l/2; i++ {
		if w[i] != w[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
	fmt.Println(firstPalindrome([]string{"notapalindrome", "racecar"}))
	fmt.Println(firstPalindrome([]string{"def", "ghi"}))
	fmt.Println(firstPalindrome([]string{"a", "b"}))
}
