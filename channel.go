//Unbuffered Channel -> Synchronous Communication
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := make(chan int)
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go PrintCounts("A", count)
	go PrintCounts("B", count)
	count <- 1
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("End Goroutines")

}

func PrintCounts(label string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel Closed")
			return
		}
		fmt.Printf("label %s receives Count: %d\n", label, val)
		if val == 10 {
			fmt.Printf("label %s Closed the channel\n", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
