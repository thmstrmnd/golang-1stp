package proofs

import "fmt"

func FlawedFloat64Two() string {
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)

	return "Flawed. The answer should be exactly 10."
}
