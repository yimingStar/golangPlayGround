package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// @Example 3
// building a signal channel -> it's use to ping our function
logCh := make(chan char, 50)
signalDoneCh := make(chan struct{})

// how to deal with implicit close timespot using "select statement"
func main() {
	// using make function to create channel
	ch := make(chan int, 50) // create a channel passing message type int

	logCh <- "send something"
	logCh <- "send something"

	signalDoneCh <- struct{}{}
}

func logger() {
	for {
		// blocking select statement
		select {
		case entry := <-logCh:
			fmt.Printf("receive, do something")
		case <- signalDoneCh:
			break;
		}
	}
}

// @Example 2
// continue reading message in the channel by for loop
func main() {
	// using make function to create channel
	ch := make(chan int, 50) // create a channel passing message type int
	wg.Add(2)
	// go func(ch <-chan int) {
	// 	// we need a close() to inform this loop need to close end exit the go routine
	// 	for i := range ch {
	// 		fmt.Printf("receive %d from channel %v\n", i, ch)
	// 	}
	// 	wg.Done()
	// }(ch)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("receive %d from channel %v\n", i, ch)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	// this is a sender, the input limit this routine only can be a sender
	go func(ch chan<- int) {
		ch <- 28
		ch <- 33
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

// @Example 1
// simple example
func main() {
	// using make function to create channel
	ch := make(chan int) // create a channel passing message type int
	wg.Add(2)
	// this is a receiver, receive int from ch
	go func() {
		i := <-ch
		fmt.Printf("receive %d from channel %v", i, ch)
		wg.Done()
	}()

	// this is a sender
	go func() {
		ch <- 28
		wg.Done()
	}()
	wg.Wait()
}
