package main

import "fmt"

func main() {
	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speed = 299_792_458 // speed of light in m / s
	)

	const time = distance / speed // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)

}
