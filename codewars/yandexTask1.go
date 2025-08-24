package codewars

import (
	"fmt"
	"log"
)


func SumYa() {
    var number int
    _, err := fmt.Scan(&number)
    if err != nil {
        log.Fatal(err)
    }
   str := sum(5, 1)
	fmt.Println(str)
	// for i := 1; i <= number; i++ {
	// 	str := sum(i, number, []int{})
	// 	fmt.Println(str)
	// }
}

func sum(number int, currentNumber int) []int {
   tempCount := 1
   var resultSlice []int = []int{currentNumber}
   var tempCurrNumb int 
   for  {
      tempCurrNumb += tempCount
      if(tempCurrNumb == number) {
          break
      }
      resultSlice = append(resultSlice, tempCurrNumb)
   }
   return resultSlice
}

// 5 
// 1 + 1 + 1 + 1 + 1
// 2 + 1 + 1 + 1
// 3 + 1 + 1
// 3 + 2
// 4 + 1