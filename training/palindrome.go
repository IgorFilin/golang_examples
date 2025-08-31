package training

import (
	"fmt"
	"unicode"
)

func Patterns() {
  result := isPalindrome("топот")
  fmt.Println(result)

}

func isPalindrome(s string) bool {
	var clearedArray []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			clearedArray = append(clearedArray, unicode.ToLower(r))
		}
	}

	n := len(clearedArray)
	for i := 0; i < n/2; i++ {
		if clearedArray[i] != clearedArray[n-1-i] {
			return false
		}
	}

	return true
}
