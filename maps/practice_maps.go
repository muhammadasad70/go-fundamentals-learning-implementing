package main

import "fmt"

func mapCreation() {
	names := map[int]string{
		1: "Muhammad",
		2: "Asad",
		3: "Hassan Nawaz",
		4: "Osman",
	}
	// name := student[1]
	// fmt.Println(student[3])
	// fmt.Println(student[10])

	// we are using the value and ok technique

	name, ok := names[4]
	// fmt.Println(name)
	// fmt.Println(ok)

	// we can also use it with the condition like
	if !ok {
		fmt.Println("Name not found")
	} else {
		fmt.Println("Name found :", name)
	}
}
