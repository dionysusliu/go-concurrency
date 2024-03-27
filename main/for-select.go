package main

import "fmt"
import "time"

// pattern 1: iterate through some iterable via channel
func forSelect_Iter(iterable []int, done <-chan bool) {
	for _, s := range iterable { // something iterable
		select {
		case <-done: // asynchronously take a signal
			return
		default:
			fmt.Printf("%d\n", s)
		}
	}
}

// pattern 2: Loop infinitely, until be signaled to stop
func forSelect_asyncStop(done <-chan bool) {
	i := 0
	for { // infinite loop
		select {
		case <-done:
			return
		default: // jump out the select to next work
		}
		// some non-preemptable work
		fmt.Println(i)
		i++
	}
}

func forSelect_main() {
	done := make(chan bool)
	fmt.Println("Main thread sleep for 1 second.")
	go forSelect_asyncStop(done)
	time.Sleep(time.Second)
	done <- true // signal the goroutine to stop
	fmt.Println("Main thread signals to stop")
}
