package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
func main () {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i :=0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v;
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}