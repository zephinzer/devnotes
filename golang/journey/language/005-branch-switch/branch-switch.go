package main

import (
	fmt "fmt"
)

func main() {
	a, b, c := 1, 2, 2

	//
	// switch with no expression
	//
	switch { // note this as compared to switch block below
	case a == 1:
		fmt.Println("a == 1")
	case b == 2:
		fmt.Println("b == 2")
	}

	switch { // first truthy evaluation is printed
	case b == 2:
		fmt.Println("b == 2")
	case a == 1:
		fmt.Println("a == 1")
	}

	//
	// switch with a variable
	//
	switch c {
	case 3:
		fmt.Println("c == 3")
	case 2:
		fmt.Println("c == 2")
	case 1:
		fmt.Println("c == 1")
	}

	//
	// switch with an expression and variable
	//
	switch d := a + c + b; d {
	case 5:
		fmt.Println("d, its 5")
	case 4:
		fmt.Println("d, its 4")
	case 3:
		fmt.Println("d, its 3")
	case 2:
		fmt.Println("d, its 2")
	case 1:
		fmt.Println("d, its 1")
	}
}
