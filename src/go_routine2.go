package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("From Goroutine 1:", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			fmt.Println("From Goroutine 2:", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutines have completed.")
}

