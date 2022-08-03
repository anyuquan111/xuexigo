package main

import (
	"fmt"
	"time"
)

func work(done chan bool) {
	fmt.Print("work.....")
	time.Sleep(5 * time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go work(done)
	<-done
}
