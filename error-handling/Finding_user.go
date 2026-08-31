package main

import (
	"errors"
	"fmt"
)

func Find_user(id int) error {
	if id <= 0 {
		return errors.New("Invalid user id")
	}
	return nil
}
func User() {
	err := Find_user(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)

}
