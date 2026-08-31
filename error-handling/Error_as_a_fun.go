package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can not divide by zero")
	}
	result := a / b

	return result, nil
}

func Handle_Error() {
	result, err := Divide(10, 0)

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result is :", result)
}
