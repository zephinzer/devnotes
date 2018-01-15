package main

import "fmt"

type NestedNumber struct {
	nested *[]NestedNumber
	value  int
}

func main() {
	var anArrayOfSize10 [10]int
	anArrayOfSize10[3] = 10
	anArrayOfSize10[8] = 11
	fmt.Println(anArrayOfSize10)

	var anotherArrayOfSize4 [4]string
	fmt.Println(anotherArrayOfSize4)

	fmt.Println("[1, [2, 3], [4, [5, [6], 7]], [[8], 9], 10]")
	arbitrarilyNestedArray := []NestedNumber{
		{value: 1},
		{nested: &[]NestedNumber{
			{value: 2},
			{value: 3},
		}},
		{nested: &[]NestedNumber{
			{value: 4},
			{nested: &[]NestedNumber{
				{value: 5},
				{nested: &[]NestedNumber{
					{value: 6},
				}},
				{value: 7},
			}}}},
		{nested: &[]NestedNumber{
			{nested: &[]NestedNumber{
				{value: 8},
			}},
			{value: 9},
		}},
		{value: 10},
	}
	fmt.Println(arbitrarilyNestedArray)
	preOrderPrint(arbitrarilyNestedArray)
	fmt.Println("\n---END---")
}

func preOrderPrint(nestedNumber []NestedNumber) {
	for i := 0; i < len(nestedNumber); {
		if nestedNumber[i].value == 0 {
			preOrderPrint(*nestedNumber[i].nested)
		} else {
			fmt.Print(nestedNumber[i].value, " ")
		}
		i++
	}
}
