package main

import (
	"fmt"
	"sync"
)

func main() {
	a := func(s string, wg *sync.WaitGroup) {
		go func() {
			fmt.Println(s)
			wg.Done()

		}()

	}
	var wg sync.WaitGroup
	wg.Add(1)
	go a("this is im passing", &wg)
	wg.Wait()
}
