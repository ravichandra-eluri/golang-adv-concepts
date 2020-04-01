package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Run with: go run -race ./racecondition/
// The race detector will flag concurrent reads/writes to incrementer.
func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer    // read (no lock)
			runtime.Gosched()   // yield so another goroutine can run
			v++
			incrementer = v     // write (no lock) — DATA RACE
			fmt.Println(incrementer)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value (unpredictable due to race):", incrementer)
}
