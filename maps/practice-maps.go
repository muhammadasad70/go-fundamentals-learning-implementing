package main

import "fmt"

func mapCreation() {
	student := map[int]string{
		1: "Muhammad",
		2: "Asad",
		3: "Hassan Nawaz",
		4: "Osman",
	}
	fmt.Println(student[3])
}
