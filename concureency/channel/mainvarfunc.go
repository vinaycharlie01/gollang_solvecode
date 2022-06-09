package main

import "fmt"

func main() {
	// a := func() <-chan int {
	// 	ch := make(chan int)
	// 	go func() {
	// 		defer close(ch)
	// 		for i := 0; i < 6; i++ {
	// 			ch <- i

	// 		}

	// 	}()
	// 	return ch

	// }
	a := func(ch chan<- int) {
		//defer close(ch)

		for i := 0; i < 6; i++ {
			ch <- i

		}
		//close(ch)
	}
	b := func(ch <-chan int) {
		for v := range ch {
			fmt.Println(v)

		}
		fmt.Println("dome..!")

	}
	ch := make(chan int)
	go a(ch)
	b(ch)

}

// package main

// import "fmt"

// func main() {

// 	owner := func() <-chan int {
// 		ch := make(chan int)

// 		go func() {
// 			defer close(ch)
// 			for i := 0; i < 5; i++ {
// 				ch <- i

// 			}
// 		}()
// 		return ch
// 	}
// 	consumer := func(ch <-chan int) {
// 		for v := range ch {
// 			fmt.Printf("Reciver: %d\n", v)
// 		}
// 		fmt.Println("Done reciving")
// 	}

// 	ch := owner()
// 	consumer(ch)
// }
