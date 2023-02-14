### Introduction

One of the most exciting characters of GO langauge is its parallel running to finish the task in a very high speed. 

In this blog, I will summary two examples to explain how to use this character.

### Example1

```
package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

```

In this example, two Goroutines are created, each with a slice of the original slice of integers. 

The Goroutines then concurrently calculate the sum of their respective slices and send the result to a common Channel c. 

The main Goroutine then receives the results from the two Goroutines and adds them together to get the final sum. 

This program demonstrates how Goroutines and Channels can be used to divide a task into smaller sub-tasks 

and coordinate their execution in parallel.


### Example2

Here is another example of a concurrent program in Go that prints the numbers from 1 to 10 using two separate goroutines:

```
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

```

In this program, two goroutines are created using anonymous functions. 

Each goroutine runs in parallel with the main goroutine and prints the numbers from 1 to 5 and 6 to 10, respectively. 

The sync.WaitGroup is used to wait for both goroutines to complete before the main goroutine terminates.

### Summary

In this blog, I summary two examples of GO concurrent program. We can create concurrent program by using `Goroutines` and `sync.WaitGroup`
