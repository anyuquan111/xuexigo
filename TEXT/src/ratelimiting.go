package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)
	for request := range requests {
		a := <-limiter
		fmt.Println("request: ", request, a, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for burstyRequest := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request :", burstyRequest, time.Now())
	}

}
