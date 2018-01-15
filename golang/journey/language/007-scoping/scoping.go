package main

import (
	fmt "fmt"
)

func main() {
	a := 5
	b := 3
	{
		b := 6
		fmt.Printf("1 --- a == %d, b == %d\n", a, b)
	}
	fmt.Printf("2 --- a == %d, b == %d\n", a, b)
}
