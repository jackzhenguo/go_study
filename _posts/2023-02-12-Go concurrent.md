Today blog I will talk about go concurrent. 

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

### Challenges

Although Go is designed to make concurrency easier, concurrent programming can still be challenging in any language. 

Here are some of the difficulties I have thought for Go language concurrent programming:

1. Concurrency bugs: Concurrent programs can have race conditions, deadlocks. And concurrency bugs are usually difficult to detect and debug. Go provides tools like `channels`, `mutexes`, and the `go` keyword to help prevent these issues, but it's still possible to write concurrent code that has bugs.

2. Synchronization and communication: In a concurrent program, it's important to synchronize access to shared data and communicate between concurrent processes. This can be challenging to get right, especially when multiple processes need to access the same data at the same time.

3. Performance: Concurrent programming can have performance issues, particularly if there is contention for shared resources like locks or channels. It's important to design concurrent programs carefully to minimize contention and maximize parallelism.


Overall, concurrent programming in Go can be challenging, but with careful design and attention to detail, it's possible to write efficient, scalable, and reliable concurrent programs in Go.



### Summary

In this blog, I summary two examples of GO concurrent program. We can create concurrent program by using `Goroutines` and `sync.WaitGroup`
