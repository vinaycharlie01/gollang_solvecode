package main

import "fmt"

func sending(ch1 chan<- string) {
	ch1 <- "hey vinay how are you...!"

}

func receving(ch1 <-chan string, ch2 chan<- string) {
	fmt.Println(<-ch1)
	m := "hey im fine how are you...dear!"
	ch2 <- m

}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sending(ch1)
	go receving(ch1, ch2)
	v := <-ch2
	fmt.Println(v)

}
