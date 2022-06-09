package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "hello hi karthil"

	}()

	select {
	case m := <-ch1:
		fmt.Println("this is first:", m)
	case <-time.After(1 * time.Second):
		fmt.Println("TIME out:")

	}

}
