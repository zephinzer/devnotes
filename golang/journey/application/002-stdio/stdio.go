package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	aString,
		aChar,
		anInteger,
		aFloat,
		aBoolean := "aString", 'a', 5, 3.1415926536, true
	fmt.Print("Print %s %c %d %f %t ", aString, aChar, anInteger, aFloat, aBoolean)
	fmt.Printf("\n")
	fmt.Printf("Printf %s %c %d %f %t ", aString, aChar, anInteger, aFloat, aBoolean)
	fmt.Printf("\n")
	fmt.Println("Println %s %c %d %f %t ", aString, aChar, anInteger, aFloat, aBoolean)
	fmt.Printf("\n")
	var someInteger int
	fmt.Printf("Enter an integer to replicate: ")
	fmt.Scanf("%d", &someInteger)
	fmt.Printf("%d\n", someInteger)
	fmt.Printf("\n")
	var someString string
	fmt.Printf("Enter a string to replicate (no spaces): ")
	fmt.Scanf("%s", &someString)
	fmt.Println(someString)
	fmt.Printf("\n")
	fmt.Printf("Enter a string to replicate (with spaces): ")
	reader := bufio.NewReader(os.Stdin)
	someOtherString, _ := reader.ReadString('\n')
	fmt.Printf("%s", someOtherString)
}
