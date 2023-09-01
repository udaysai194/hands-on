package main

import (
	"fmt"
	"sync"
	"time"
)

// func go_routine() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		i := i
// 		go func() {
// 			greeter("hello", i)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()

// }

// func greeter(s string, i int) {
// 	fmt.Println(s, i)
// }

var wg sync.WaitGroup

func go_routine() {
	go greeter("hello", 1)
	go greeter("world", 2)
	wg.Add(2)
	wg.Wait()

}

func greeter(s string, i int) {
	if s == "world" {
		time.Sleep(3 * time.Second)
	}
	fmt.Println(s, i)
	wg.Done()
}
