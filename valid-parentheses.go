package main

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	bracketsPairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	unclosedBrackets := ""

	for _, char := range s {
		beginningBracket, ok := bracketsPairs[char]
		if ok {
			if len(unclosedBrackets) == 0 {
				return false
			}
			// check the beginningBracket is the last one in the stack
			lastInStack := unclosedBrackets[len(unclosedBrackets)-1:]
			if string(beginningBracket) != lastInStack {
				return false
			}
			//remove the last rune from the string
			unclosedBrackets = unclosedBrackets[:len(unclosedBrackets)-1]
		} else {
			// add the bracket to the stack
			unclosedBrackets += string(char)
		}
	}

	if len(unclosedBrackets) > 0 {
		return false
	}

	return true
}

//func main() {
//	//fmt.Printf("Result: %t", isValid("()"))
//	//fmt.Printf("Result: %t", isValid("()[]{}"))
//	//fmt.Printf("Result: %t", isValid("(]"))
//	//fmt.Printf("Result: %t", isValid("["))
//	fmt.Printf("Result: %t", isValid("]"))
//}
