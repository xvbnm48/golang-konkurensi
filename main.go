package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 + time.Millisecond)
		fmt.Printf(s)
	}
}
func main() {
	go say("world")
}