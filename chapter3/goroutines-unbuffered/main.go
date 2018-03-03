package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	count := make(chan int)

	wg.Add(4)

	fmt.Println("Start Goroutines")

	go printCounts("A", count)
	go printCounts("B", count)
	go printCounts("C", count)
	go printCounts("D", count)

	fmt.Println("Channel begin")
	count <- 1

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts(label string, count chan int) {
	defer wg.Done()

	for {
		// Receives message from channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed", ok)
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel closed from %s \n", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
