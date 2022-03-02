package utils

import (
	"fmt"
	"strconv"
)

func ConvertToString(n interface{}) string {
	return fmt.Sprintf("%v", n)
}

func IsTrue(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
