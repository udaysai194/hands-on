package main

import (
	"fmt"
	"sync"
)

func mutexes() {

	wg := &sync.WaitGroup{}
	semaphore := &sync.Mutex{}

	sharedMemory := []int{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, semaphore *sync.Mutex) {
		semaphore.Lock()
		sharedMemory = append(sharedMemory, 1)
		semaphore.Unlock()
		wg.Done()
	}(wg, semaphore)

	go func(wg *sync.WaitGroup, semaphore *sync.Mutex) {
		semaphore.Lock()
		sharedMemory = append(sharedMemory, 2)
		semaphore.Unlock()
		wg.Done()
	}(wg, semaphore)

	go func(wg *sync.WaitGroup, semaphore *sync.Mutex) {
		semaphore.Lock()
		sharedMemory = append(sharedMemory, 3)
		semaphore.Unlock()
		wg.Done()
	}(wg, semaphore)

	wg.Wait()
	fmt.Println(sharedMemory)

}
