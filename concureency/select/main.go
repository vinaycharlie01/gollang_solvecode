package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hello hi karthil"

	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hello hi vinay"

	}()
	for i := 0; i < 2; i++ {
		select {
		case m := <-ch1:
			fmt.Println("this is first:", m)
		case m := <-ch2:
			fmt.Println("this is second:", m)

		}

	}

}
