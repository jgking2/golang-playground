package main

import (
	"fmt"
	"playground/concurrency"
	"playground/functions"
	"playground/lists"
	"playground/logic"
	"playground/math"
	"playground/pointers"
)

func main() {
	concurrency.ExploreConcurrencyWithMultipleChannels()
	// concurrency.ExploreConcurrencyWithChannels()
	// concurrency.ExploreConcurrency()
	pointers.AllThingsPointers()
	lists.AllThingsSlices()
	//Illustrates that anything with uppercase variable name is exported
	fmt.Println(functions.PublicConst)
	test := functions.NakedReturns(5)
	fmt.Println(test)
	fmt.Println("Whattup golang! ", math.Square(2))
	err := functions.ErrorReturn(true)
	if err != nil {
		fmt.Println(err)
	}
	first, second, third := functions.MultiReturn()
	fmt.Printf("%g %g %g", first, second, third)

	isHotDog := logic.SwitchWithInit("Hot Dog")
	fmt.Println(isHotDog)
}
