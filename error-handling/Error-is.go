package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("User Not Found")

func Get_user(id int) error {
	if id == 10 {
		return fmt.Errorf("get user %d: %w", id, ErrUserNotFound)
	}
	return nil
}
func Error_Is_implement() {
	err := Get_user(10)

	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("User does not exist:")
	}
}
