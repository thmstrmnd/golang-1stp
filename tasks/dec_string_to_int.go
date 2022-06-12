package tasks

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func DecNumStringToInt() string {
	var amount string = "2.50"
	amtHasDecPoint := strings.ContainsAny(amount, ".")

	if amtHasDecPoint {
		numStr := strings.Split(amount, ".")
		whole, frac := numStr[0], numStr[1]
		lengthDecPlaces := len(frac)

		amountWholeInt, _ := strconv.Atoi(whole)
		amountFracInt, _ := strconv.Atoi(frac)

		if lengthDecPlaces == 1 {
			amountInt := amountWholeInt*100 + amountFracInt*10

			fmt.Println(whole, reflect.TypeOf(whole))
			fmt.Println(frac, reflect.TypeOf(frac))
			fmt.Println(amountInt, reflect.TypeOf(amountInt))

		} else if lengthDecPlaces == 2 {
			fracSplit := strings.Split(frac, "")
			// fmt.Println(fracSplit)
			leftmostDigit, rightmostDigit := fracSplit[0], fracSplit[1]
			// fmt.Println("leftmostDec:", leftmostDec, "leftmostDec:", rightmostDec)

			tenths, _ := strconv.Atoi(leftmostDigit)
			hundredths, _ := strconv.Atoi(rightmostDigit)

			amountInt := amountWholeInt*100 + tenths*10 + hundredths

			fmt.Println(whole, reflect.TypeOf(whole))
			fmt.Println(frac, reflect.TypeOf(frac))
			fmt.Println(amountInt, reflect.TypeOf(amountInt))
		}
	} else {
		amountInt, _ := strconv.Atoi(amount)
		fmt.Println(amountInt, reflect.TypeOf(amountInt))
	}
	return "---"
}
