package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go countChan("sheep", c)

	for msg := range c{
		fmt.Println(msg)
	}
}

func countChan(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
