package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func sendData(channel chan string) {
  defer wg.Done()
  channel <- "Первая запись в канал"
  time.Sleep(1 * time.Second)
  channel <- "Последняя запись в канале после одной секунды"
}


func main() { 
  wg.Add(1)
  channel := make(chan string, 2)
  go sendData(channel)
  fmt.Println("Ожидание получение первых данных из канала")
  fmt.Println(<-channel)
  fmt.Println("Ожидание получение вторых данных из канала")
  fmt.Println(<-channel)
  fmt.Println("Перед ожидание завершение дочернего потока")
  wg.Wait()
  fmt.Println("Завершается основной поток")
}

// ---- Аннотация
// Каналы используются для обмена данными между горутинами. Они обеспечивают синхронизацию и безопасность.