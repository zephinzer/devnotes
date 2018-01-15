package main

import "fmt"

type CustomType1 struct {
	aString string
	aNumber int
}

type CustomType2 struct {
	x, y, z int
}

type CustomType3 struct {
	x, y, z string
}

type RecursiveType struct {
	recursiveType *RecursiveType
	value         int
}

func main() {
	customType1 := CustomType1{"hi", 5}
	fmt.Println("customType1.1", customType1)

	customType1.aString = "bye"
	fmt.Println("customType1.2", customType1)

	customTypeOne := &customType1
	customTypeOne.aString = "hello"
	fmt.Println("customType1.3", customType1)

	customType2 := CustomType2{}
	fmt.Println("customType2", customType2)

	customType3 := CustomType3{}
	fmt.Println("customType3", customType3)

	customType3dot1 := CustomType3{"1", "", "2"}
	fmt.Println("customType3dot1", customType3dot1)

	recursiveType := RecursiveType{
		recursiveType: &RecursiveType{
			recursiveType: &RecursiveType{
				value: 4,
			},
			value: 5,
		},
		value: 6,
	}
	fmt.Println(recursiveType)
}
