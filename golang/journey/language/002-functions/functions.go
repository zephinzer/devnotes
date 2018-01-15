package main

import (
	fmt "fmt"

	"./diffPackage"
	renamedPackage "./diffRefPackage"
)

func main() {
	defer (func() {
		fmt.Println("HELLO AFTER THE PROGRAM RETURNS!")
	})()
	helloWorld()
	sayHelloTo("joseph")
	helloFromTheSamePackage()
	diffPackage.HelloFromADifferentPackage()
	renamedPackage.HelloFromARenamedPackage()
	variableArguments("hi")
	variableArguments("hi2", 1, 2, 3, 4)
	aClosure := closure("closure: nope, print this instead.\n")
	aClosure()
	fmt.Print("the function is located at memaddr[", closure("print this!!!!"), "]\n")
}

func closure(printMe string) func() int {
	return func() int {
		fmt.Print(printMe)
		return 1
	}
}

func variableArguments(base string, intArgs ...int) {
	fmt.Print(base, intArgs, "\n")
}

func helloWorld() {
	fmt.Printf("Hello world!\n")
}

func sayHelloTo(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
