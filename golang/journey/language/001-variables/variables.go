package main

import "fmt"

func main() {
	fmt.Printf("\nTEXTUAL")
	fmt.Printf("\n-------\n")

	var someString = "Hello world!"
	fmt.Printf("string:                  %s\n", someString)

	var someChar = 'a'
	fmt.Printf("char:                    %c\n", someChar)

	var someTypecastedChar int = 'a'
	fmt.Printf("typecasted char:         %d\n", someTypecastedChar)

	var explicitString string = "Hello world."
	fmt.Printf("explicit string:         %s\n", explicitString)

	fmt.Printf("\nNUMBERS")
	fmt.Printf("\n-------\n")

	var someInteger = 5
	fmt.Printf("integer:                 %d\n", someInteger)

	var someTypecastedInteger byte = 97
	fmt.Printf("typecasted to char:      %c\n", someTypecastedInteger)

	var explicitInteger int = 6
	fmt.Printf("explicit integer:        explicitInteger == %d\n", explicitInteger)

	var negativeInteger int = -7
	fmt.Printf("negative integer:        negativeInteger == %d\n", negativeInteger)

	var someFloat = 3.1415926536
	fmt.Printf("float:                   someFloat == %f\n", someFloat)

	var explicitFloat float64 = 1.618033988749894848
	fmt.Printf("explicit float:          explicitFloat == %f\n", explicitFloat)

	var boolean = true
	fmt.Printf("boolean:                 %t\n", boolean)

	fmt.Printf("\nDECLARATIONS")
	fmt.Printf("\n------------\n")

	a := 5
	fmt.Printf("lazy declare:            a == %d\n", a)

	b, c := 3, 4
	fmt.Printf("mutli-assign:            b == %d, c == %d\n", b, c)

	d, e := "hello", 8
	fmt.Printf("multi-assign-type:       d == %s, e == %d\n", d, e)

	fmt.Printf("\nFORMATTING")
	fmt.Printf("\n----------\n")

	aString,
		aChar,
		anInteger,
		aFloat,
		aBool := "Hello", 'A', 42, 3.142, true

	fmt.Printf("all in one line:         %s %c %d %f %t\n", aString, aChar, anInteger, aFloat, aBool)

	fmt.Printf("aString typecheck:       %T\n", aString)
	fmt.Printf("aChar typecheck:         %T\n", aChar)
	fmt.Printf("anInteger typecheck:     %T\n", anInteger)
	fmt.Printf("aFloat typecheck:        %T\n", aFloat)
	fmt.Printf("aBool typecheck:         %T\n", aBool)

	fmt.Printf("anInteger in binary:     %b\n", anInteger)
	fmt.Printf("anInteger in octal:      %#o\n", anInteger)
	fmt.Printf("anInteger in hex:        %#x\n", anInteger)
	fmt.Printf("anInteger in HEX:        %#X\n", anInteger)

	fmt.Printf("integer formatting:      explicitInteger == %+d\n", explicitInteger)
	fmt.Printf("ninteger formatting:     negativeInteger == %+d\n", negativeInteger)
	fmt.Printf("float formatting:        someFloat == %.32f\n", someFloat)
	fmt.Printf("float formatting:        explicitFloat == %.32f\n", explicitFloat)

	fmt.Printf("reminder that a:         a == %d", a)
	fmt.Printf("memory address of a:     %#x\n", &a)
	f := a
	fmt.Printf("memory address of f(=a): %#x\n", &f)
	g := a
	fmt.Printf("memory address of g(=a): %#x\n", &g)
	h := &a
	fmt.Printf("value of h(=&a):         %d\n", *h)
}
