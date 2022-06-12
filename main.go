package main

import (
	"1stp-golang/proofs"
	"fmt"
)

func main() {
	// tasks := tasks.DecNumStringToInt()
	// fmt.Println(tasks)

	// proofOne := proofs.FlawedFloat64One()
	// fmt.Println(proofOne)

	// proofTwo := proofs.FlawedFloat64Two()
	// fmt.Println(proofTwo)

	proofThree := proofs.FlawedFloat64Three()
	fmt.Println(proofThree)
}
