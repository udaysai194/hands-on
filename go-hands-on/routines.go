package main

import (
	"fmt"
	"time"
)

func go_routine() {
	go greeter("hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
