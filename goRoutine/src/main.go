package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

// RW mutex allow mutiple read but only one write
var m = sync.RWMutex{}
var counter int = 0

func main() {
	fmt.Printf("Thread numbers %v\n", runtime.GOMAXPROCS(-1))
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go printCounter()
		m.Lock()
		go increaseCounter()
	}
	wg.Wait()
}

func printCounter() {
	fmt.Printf("The counter: %d\n", counter)
	m.RUnlock()
	wg.Done()
}

func increaseCounter() {
	counter++
	m.Unlock()
	wg.Done()
}
