package main

import (
	"1stp-golang/tasks"
	"fmt"
)

func main() {
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

	ratNum := tasks.BigRational()
	fmt.Println(ratNum)

}
