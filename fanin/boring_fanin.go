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
	c := fanIn(boring("joe"), boring("Anne"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You're boring: I'm Leaving")
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
