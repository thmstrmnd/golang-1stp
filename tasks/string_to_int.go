package Tasks

import (
	"fmt"
	"reflect"
	"strings"
)

func StringToInt() string {
	numStr := strings.Split("3.0", ".")
	whole, decimal := numStr[0], numStr[1]
	testNum := 3.5

	fmt.Println(whole, reflect.TypeOf(whole))
	fmt.Println(decimal, reflect.TypeOf(decimal))
	fmt.Println(testNum, reflect.TypeOf(testNum))

	return "---"
}
