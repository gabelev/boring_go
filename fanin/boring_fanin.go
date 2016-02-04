package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn(input, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := boring("gabe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow")
			return
		}
	}
}

func boring(msg string) <-chan Message { // returns receive-only channel of strings
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() { // we launch goroutine from inside func
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // returns the channel to the caller
}
