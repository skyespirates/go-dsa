package main

import (
	"fmt"
	"unicode"
)

func isAlphaNumeric(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char <= '9'
}

func isPalindrome(s string) bool {

	ch := []rune(s)

	i, j := 0, len(ch)-1

	for i < j {
		if !isAlphaNumeric(ch[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(ch[j]) {
			j--
			continue
		}

		if unicode.ToLower(ch[i]) != unicode.ToLower(ch[j]) {
			return false
		}
		i++
		j--
	}

	return true

}

func main() {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	fmt.Println(result)
}
