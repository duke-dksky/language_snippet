package main

import (
	"fmt"
	"sync"
)

// first stage
func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// second stage
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// fan-in
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-done:
					return
				}
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// finnal stage
func main() {

	done := make(chan struct{})
	defer close(done)

	c := gen(done, 2, 3)
	c1, c2 := sq(done, c), sq(done, c)

	out := merge(done, c1, c2)
	fmt.Println(<-out)
	//for n := range merge(out1, out2) {
	//	fmt.Println(n)
	//}

}
