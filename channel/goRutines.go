package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Messages 2")
	go f("Hello world")
	time.Sleep(5 * time.Second)
}
