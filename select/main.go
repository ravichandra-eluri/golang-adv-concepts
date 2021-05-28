package main

import (
	"fmt"
	"time"
)

func main() {
	// Select: block until one of the cases is ready
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "one"
	ch2 <- "two"

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("ch1:", msg)
		case msg := <-ch2:
			fmt.Println("ch2:", msg)
		}
	}

	// Default case: non-blocking receive
	ch := make(chan int, 1)
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("no value ready (non-blocking)")
	}

	// Timeout: give up if no value arrives in time
	slow := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "slow result"
	}()

	select {
	case result := <-slow:
		fmt.Println("got:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out waiting for slow channel")
	}

	// Done signal: break out of a channel-drain loop
	nums := make(chan int, 5)
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		nums <- i
	}
	close(nums)
	go func() {
		time.Sleep(10 * time.Millisecond)
		close(done)
	}()

loop:
	for {
		select {
		case n, ok := <-nums:
			if !ok {
				fmt.Println("nums channel closed")
				break loop
			}
			fmt.Println("num:", n)
		case <-done:
			fmt.Println("done signal received")
			break loop
		}
	}
}
