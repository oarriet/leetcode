package main

import (
	"strings"
)

/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")
	return len(split[len(split)-1])
}

//func main() {
//	fmt.Println(lengthOfLastWord("Hello World"))
//	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
//	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
//}
