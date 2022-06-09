package main

import "fmt"

//this is strnger
// this is sending messages to vinay
func foo(ch chan<- string) {
	b := "Hi  Vinay im stranger how are you...! "
	ch <- b

}
func seefoo(ch <-chan string) {
	fmt.Println(<-ch)

}
func see(ch <-chan string, c chan<- string) {
	fmt.Println(<-ch)
	d := "hello Sranger how are you...!"
	c <- d
}
func foosee(c <-chan string) {
	fmt.Println(<-c)

}

func main() {
	ch := make(chan string)
	c := make(chan string)
	go foo(ch, c)
	seefoo(ch)
	go see(ch, c)
	seefoo(c)

}
