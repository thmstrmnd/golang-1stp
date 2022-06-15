package main

import (
	"1stp-golang/poc"
	"fmt"
)

type Border interface {
	Perimeter() float64
}

func printBorder(b Border) {
	fmt.Println(b.Perimeter())
}

func main() {
	c := poc.Circle{5}
	r := poc.Rectangle{20, 10}

	printBorder(c)
	printBorder(r)

	// tasks := tasks.DecNumStringToInt()
	// fmt.Println(tasks)

	// tasks := tasks.IntCap()
	// fmt.Println(tasks)

	// proofOne := poc.FlawedFloat64One()
	// fmt.Println(proofOne)

	// proofTwo := poc.FlawedFloat64Two()
	// fmt.Println(proofTwo)

	// proofThree := poc.FlawedFloat64Three()
	// fmt.Println(proofThree)

	// stringFromInt64 := tasks.Int64ToString()
	// fmt.Println(stringFromInt64)

	// ratNum := tasks.BigRational()
	// fmt.Println(ratNum)

	// methodsPoc1 := poc.MethodsPoc()
	// fmt.Println(methodsPoc1)

}
