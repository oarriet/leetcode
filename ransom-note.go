package main

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*/

func canConstruct(ransomNote string, magazine string) bool {

	var magazineLetters = map[string]int{}

	for _, letter := range magazine {
		magazineLetters[string(letter)] = magazineLetters[string(letter)] + 1
	}

	for _, letter := range ransomNote {
		if value, ok := magazineLetters[string(letter)]; ok && value > 0 {
			magazineLetters[string(letter)] = value - 1
		} else {
			return false
		}
	}

	return true
}

//func main() {
//	//println(canConstruct("a", "b"))
//	//println(canConstruct("aa", "ab"))
//	//println(canConstruct("aa", "aab"))
//
//	//marshal, err := json.Marshal(nil)
//	//if err != nil {
//	//	println("error")
//	//	println(err)
//	//}
//	//println("marshal")
//	//println(string(marshal))
//	//println("end marshal")
//
//	var versionRangeTupleList [][]string
//
//	err := json.Unmarshal([]byte("[]"), &versionRangeTupleList)
//	if err != nil {
//		println("error")
//		println(err)
//	} else {
//		println("success")
//		println(versionRangeTupleList)
//	}
//}
