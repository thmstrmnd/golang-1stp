package proofs

import "fmt"

func FlawedFloat64Three() string {
	var amount1 float64 = 10
	var amount2 float64 = 0.01255876596569785670685694569576590785

	// fmt.Println(len(amount2))

	mulAmount := amount1 * amount2
	fmt.Println(mulAmount)

	return "---"
}

// Reference
// https://www.educative.io/edpresso/what-is-type-float64-in-golang

// Notes
// mulAmount result does not reflect the correct cound of decimal places.
