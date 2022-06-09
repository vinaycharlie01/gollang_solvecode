package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)

	}

}

func main() {
	//this is normal
	fun("hello go")

	//this os by using go rotine
	go fun("this is go rotine!")
	//ananumus fuction
	go func() {
		fun("This is ananums fuction")

	}()
	fv := fun
	go fv("this is value call.....!")
	fmt.Println("this is waiting for gortine...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Done...")
}
