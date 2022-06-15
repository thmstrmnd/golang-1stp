package tasks

import (
	"fmt"
	"log"
	"math/big"
	"reflect"
)

func BigRational() string {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	r := new(big.Rat)
	_, err := fmt.Sscan("1.5000", r)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(r)
	}

	rDenom := r.Denom().Int64() // can be r.Denom().String()
	fmt.Println(rDenom, reflect.TypeOf(rDenom))

	return "---"
}
