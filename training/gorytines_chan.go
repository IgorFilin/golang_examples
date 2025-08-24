package training

import "fmt"

func Async() {
	var array = []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	go flowFunc(array, ch)

	for {
		resultCh, isExist  := <- ch
		if(!isExist) {
			break
		}
		fmt.Println(resultCh)
	}
	fmt.Println("Завершение")
}

func flowFunc(array []int, ch chan int) {
  defer close(ch)
  for _, value := range array {
    ch <- value
  }
}

// Задача сделать, что бы на вход был срез из чисел, и каждое из чисел выводилась потоком из
// параллельного потока горутины, с таймаутом в 2 секунды