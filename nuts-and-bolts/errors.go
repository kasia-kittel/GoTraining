package nutsandbolts

import (
	"errors"
	"fmt"
)

// It’s possible to use custom types as errors by implementing the Error() method on them.
type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func CantDoWith42(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func CantDoWith50(arg int) (int, error) {

	if arg == 50 {

		return -1, errors.New("can't work with it")
	}
	return arg + 3, nil
}
