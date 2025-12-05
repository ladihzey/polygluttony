package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Synchronous execution:")
	fmt.Print("[")
	for i := 0; i < 30; i++ {
		PrintEmpty()
		PrintFilled()
	}
	fmt.Print("]\n\n")

	fmt.Println("Concurrent execution:")
	fmt.Print("[")
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		RunWithWaitGroup(&wg, PrintEmpty)
		RunWithWaitGroup(&wg, PrintFilled)
	}
	wg.Wait()
	fmt.Print("]\n\n")
}

func PrintEmpty() {
	fmt.Print(" ")
}

func PrintFilled() {
	fmt.Print("▉")
}

func RunWithWaitGroup(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
