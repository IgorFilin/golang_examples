package main

import (
	"fmt"
	"unicode"
)

// "demo/app/codewars"/
// "demo/app/training"


func main() { 
  result := isPalindrome("Топотт")
  fmt.Println("Результат: ")
  fmt.Println(result)
}


// func isPalindrome(s string) bool {
// 	trimStr := strings.Trim(s, "")
// 	lowerStr := strings.ToLower(trimStr)
// 	arrayWithStr := strings.Split(lowerStr, "")
// 	slices.Reverse(arrayWithStr)
// 	fmt.Println(arrayWithStr)
// 	updatedStr := strings.Join(arrayWithStr,"")
// 	return lowerStr == updatedStr
// }


func isPalindrome(s string) bool {
	var clearedArray []rune
	
	for _, r := range s {
		if(unicode.IsLetter(r)) {
			clearedArray = append(clearedArray, unicode.ToLower(r))
		}
	}

	n := len(clearedArray)
	for i := 0; i < n/2; i++ {
        if(clearedArray[i] != clearedArray[n - 1 - i]) {
			return false
		}
	}
	
	return true
}


