package main

import (
	"fmt"
)

func F_Err(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid user id:%d", id)
	}
	return nil
}
func Formating() {
	id := -1
	if err := F_Err(id); err != nil {
		fmt.Println(err)
	}

}
