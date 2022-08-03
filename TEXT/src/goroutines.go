package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("an")
	go f("yu")
	go f("wang")
	go func() { fmt.Println("quan") }()
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
