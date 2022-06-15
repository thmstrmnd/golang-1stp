package poc

import "fmt"

func FlawedFloat64One() string {
	var amount1 float64 = 268.0
	var amount2 float64 = 0.0125

	mulAmount := amount1 * amount2
	fmt.Println(mulAmount)

	return "No flaw."
}
