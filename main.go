package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func fakeApiRequest() map[string]string  {
  time.Sleep(4 * time.Second)
  response := map[string]string{
    "name": "Igor",
  }
  return response
}

func sendData(channel chan map[string]string) {
  defer wg.Done()
  channel <- fakeApiRequest()
  fmt.Println("Завершился дополнительный поток")
}


func main() { 
  wg.Add(1)
  channel := make(chan map[string]string)
  go sendData(channel)
  for i := 0; i <= 10; i++ {
    fmt.Printf("Основной поток %d \n", i)
  }
  result := <- channel
  fmt.Print(result)
  wg.Wait()
  fmt.Println("Завершается основной поток")
}

// ---- Аннотация
// Каналы используются для обмена данными между горутинами. Они обеспечивают синхронизацию и безопасность.