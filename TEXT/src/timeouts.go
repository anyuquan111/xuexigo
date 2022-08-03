package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "result  1"
	}()
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("TIMEOUT")

	}
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case aa := <-time.After(time.Second):
		fmt.Println("timeout")
		fmt.Println(aa)
	}

}
