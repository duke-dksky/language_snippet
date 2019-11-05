package main

import (
	"fmt"
	"sync"
	"time"
)

func closedChannel() {
	finish := make(chan struct{})
	var wg sync.WaitGroup
	const n = 20
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-finish:
			case <-time.After(1 * time.Hour):
			}
		}()
	}
	t0 := time.Now()
	close(finish)
	wg.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}

func WaitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
	}
}

func main() {
	var ch chan bool
	<-ch
}
