package main

import "math"

// for _, v := range a {
// 	b := math.Max(float64(v), -0)

// }
func main() {

	b := func(a []int) {
		b := func(a []int) {
			for _, v := range a {
				b := math.Max(float64(v), -0)
				

			}

		}
		b(a)

	}
	b([]int{10, 20, 30, 40})

}
