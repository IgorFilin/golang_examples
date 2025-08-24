package training

import (
	"fmt"
	"sync"
)

var counter int

func MutexTrainig() {
	var mutex sync.Mutex 
	for i := 0; i <= 3; i++ {
		go Foo(&mutex)
	}
	fmt.Scanln()
}

func Foo(mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0

	for i := 0; i <= 10; i++ {
		counter += 1
		fmt.Println(counter)
	}
	mutex.Unlock()
}

// Закомментировать mutex.Lock(), mutex.Unlock() для просмотра как он влияет на глобальные переменные при использовании горутин