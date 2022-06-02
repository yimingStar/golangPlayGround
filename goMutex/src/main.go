package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	routineNum := 1
	wg.Add(routineNum)
	msg := "Hello!"
	go shoutOut(msg)

	wg.Wait()
}

func shoutOut(msg string) {
	fmt.Println(msg)
	wg.Done()
}

/**
func demomain() {
	msg := "Hello!"

	go func() {
		fmt.Println(msg)
	}()

	// race condition occurs, this is bad, use argument
	// this example go routine will go function stack looking for msg
	msg = "Goodbye"

	// this is not a good pratice, do not use sleep call, use waitGroup
	time.Sleep(100 * time.Millisecond)
}
*/
