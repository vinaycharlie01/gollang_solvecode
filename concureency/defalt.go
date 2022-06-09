package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {

			ch <- "Message"

		}

	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no message")

		}

	}

}
