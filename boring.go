package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe") // function returning a channel
	anne := boring("anne")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("my name is %q\n", <-anne)

	}
	fmt.Println("You're boring: I'm Leaving")
}

func boring(msg string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() { // we launch goroutine from inside func
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // returns the channel to the caller
}
