package main

import (
	"fmt"
)

type MyError struct{}

func (e MyError) Error() string {
	return "Something goes wrong"
}

func Implement_er() {
	var err error = MyError{}
	fmt.Println(err)
}
