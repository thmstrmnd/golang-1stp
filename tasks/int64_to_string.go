package tasks

import (
	"fmt"
	"reflect"
	"strconv"
)

func Int64ToString() string {
	s := strconv.Itoa(9223372036854775807)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	return "---"
}
