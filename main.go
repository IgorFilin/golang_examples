package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	// Уведомляем WaitGroup о завершении
	defer wg.Done()
	defer fmt.Println("Горутина 1 завершилась")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Горутина 1 %d \n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printNumbers2(wg *sync.WaitGroup) {
	// Уведомляем WaitGroup о завершении
	defer wg.Done()
	defer fmt.Println("Горутина 2 завершилась")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Горутина 2 %d \n", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// Создаем переменную с группой горутин
	var wg sync.WaitGroup
	// Добавляем ожидание одной гарутины, кроме основной
	wg.Add(2)
	// Запускаем гарутину и передаём ссылку на переменную горутин
	go printNumbers(&wg)
	go printNumbers2(&wg)
	for i := 0; i <= 10; i++ {
		fmt.Println("Основной поток")
	}
	// Ожидаем завершения горутины
	wg.Wait()
}

// ---- Аннотация
// Если вам нужно дождаться завершения одной горутины, используйте каналы.
// Если вам нужно дождаться завершения нескольких горутин, используйте sync.WaitGroup.
// Ключевое слово defer в Go используется для отложенного выполнения функции или метода. Когда вы пишете defer, функция, следующая за ним, будет выполнена не сразу, а перед возвратом из текущей функции.