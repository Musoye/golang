// Goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go sayHello()
	go sayWorld()
	wg.Wait()
	fmt.Println("End Goroutines")
}

func sayHello() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello")
	}
}

func sayWorld() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("World")
	}
}
