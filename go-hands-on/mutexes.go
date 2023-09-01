package main

import (
	"fmt"
	"sync"
)

func mutexes() {

	wg := &sync.WaitGroup{}
	//var semaphore *sync.Mutex

	sharedMemory := []int{}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		sharedMemory = append(sharedMemory, 1)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		sharedMemory = append(sharedMemory, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		sharedMemory = append(sharedMemory, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(sharedMemory)

}
