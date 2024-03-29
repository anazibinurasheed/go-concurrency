package main

import (
	"fmt"
	"math/rand"
	"time"
)
//Generator: function that returns a channel

//Channels are first-class values, just like strings or integers.



func boring(msg string) <-chan string { //returns receive-only channel of strings.

	c := make(chan string)

	go func() { // we launch a goroutine from inside the function

		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)

			if i == 100{
				close(c)
			}
		}

		
	}()

	return c // return the channel to the caller

}

func main() {
	c := boring("boring!") // function returns a channel.

	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
	}
	

	fmt.Println("I'm leaving.")
}
