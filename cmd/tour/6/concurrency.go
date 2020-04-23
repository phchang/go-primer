package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/* concurrency -> goroutines
	- goroutine is not necessarily an OS level thread
		- lightweight thread (could be a process or a thread -> managed by the go runtime)
	*/

	worldPrinter := func() {
		fmt.Println("world")
	}

	go worldPrinter()

	fmt.Println("hello")

	time.Sleep(100 * time.Millisecond)

	go func() {
		fmt.Println("hello world from an inline function")
	}()

	time.Sleep(100 * time.Millisecond)

	/* channels

	<- is the channel operator, indicates data flow

	someChannel <- v // send v in the channel
	<-someChannel    // receive something from the channel
	*/

	ch := make(chan int)

	sum := func(s []int, c chan int) {

		// can specify data flow direction of channel in param declaration
		//   <-chan only allows receiving from the channel
		//   chan<- only allows sending to the channel
		//   chan   allows both sending and receiving

		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // this is blocking
		// blocks until something receives from the channel
	}

	go sum([]int{1, 2, 3}, ch)

	total := <-ch // this is also blocking
	// blocks until something is sent to the channel

	// By default, sends and receives block until the other side is ready.

	fmt.Println("total = ", total)

	/** buffered channels
	- buffer length is the second param into `make`
	*/

	buffered := make(chan int, 2) // block happens only when buffer is full

	buffered <- 1
	buffered <- 2
	//buffered <- 3 // this blocks execution

	fmt.Println("from buffered channel: ", <-buffered)
	fmt.Println("from buffered channel: ", <-buffered)

	// one example use of buffered channels is to limit resource usage
	// e.g. calling a downstream connection, database connections

	/* closing channels
	- if sender tries to send something into a closed channel, panic
	- typically you won't need to worry about closing channels
	*/

	// range receives values from a channel until it is closed

	helloChannel := make(chan rune, 11)

	for _, c := range "hello world" {
		helloChannel <- c
	}
	close(helloChannel)

	for c := range helloChannel {
		fmt.Printf("%v\n", string(c))
	}

	// select
	quit := make(chan struct{})

	isDone := false

	go func() {
		count := 0
		for {
			if isDone {
				return
			}
			time.Sleep(15 * time.Millisecond)
			count++
			fmt.Printf("%v ", count)
		}
	}()

	go func() {
		time.Sleep(250 * time.Millisecond)
		quit <- struct{}{}
	}()

	select {
	case <-quit:
		fmt.Println("quitting")
		isDone = true
	}

	//for {
	//	select {
	//	case <-quit:
	//		fmt.Println("quitting")
	//		return
	//	default:
	//		fmt.Printf(".")
	//		time.Sleep(1 * time.Millisecond)
	//	}
	//}

	/* ## 27. mutex
	- mutex -> mutual exclusion
	*/

	results := make([]string, 0, 200)

	mutex := sync.Mutex{}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			mutex.Lock()
			results = append(results, fmt.Sprintf("%v", i))
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 100; i < 200; i++ {
			mutex.Lock()
			results = append(results, fmt.Sprintf("%v", i))
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("length of results = ", len(results))
}
