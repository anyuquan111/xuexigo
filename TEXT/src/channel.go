package main

import "fmt"

func main() {
	massage := make(chan string)

	go func() { massage <- "ping" }()
	msg := <-massage
	fmt.Println(msg)
}
