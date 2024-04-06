package utils

import (
	"fmt"
)

func EscalateError(err error) {
	panic(fmt.Sprintf("%v", err))
}
