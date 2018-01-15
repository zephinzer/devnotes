package main

import (
	fmt "fmt"
)

func main() {
	a, b, c := 1, 2, 2
	fmt.Printf("%d > %d == %t\n", a, b, a > b)
	if a > b {
		fmt.Printf("\ta is more than b\n")
	} else if b > c {
		fmt.Printf("\tb is more than c\n")
	} else if c > b {
		fmt.Printf("\tc is more than b")
	} else {
		fmt.Printf("\tb is more than a\n")
	}

	// usage of short statement
	if d := 10; d > a {
		fmt.Println("d is above all else")
	}
}
