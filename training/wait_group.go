package training

import (
	"fmt"
	"sync"
)

var waitGr sync.WaitGroup

func WaitGroupTraining() {
	waitGr.Add(3)
	go WaitGroupTrainingFoo(1)
	go WaitGroupTrainingFoo(2)
	go WaitGroupTrainingFoo(3)
	waitGr.Wait()
}

func WaitGroupTrainingFoo(intenger int) {
    defer waitGr.Done()
	fmt.Println(intenger)
}