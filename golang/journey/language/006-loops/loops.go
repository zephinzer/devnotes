package main

import (
	fmt "fmt"
)

func main() {
	// standard c-semantics for
	for a := 1; a < 10; a++ {
		fmt.Println(a)
	}

	fmt.Println("---")

	// usage with short statement
	startsAt, endsAt := 1, 10
	for inTheMiddle := 5; inTheMiddle < endsAt; {
		fmt.Println(inTheMiddle)
		inTheMiddle++
	}

	fmt.Println("---")

	// usage as a while loop
	for endsAt > startsAt {
		fmt.Println(startsAt)
		startsAt++
	}
}
