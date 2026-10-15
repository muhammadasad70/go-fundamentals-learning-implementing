package main

import "fmt"

func checkingTask() {
	students := map[int]string{
		1: "Asad",
		2: "Ali",
		3: "Ahmed",
	}

	for id, name := range students {
		if name == "Ahmed" {
			fmt.Println("Key is ", id)
		}
	}

	// for i := 0; i <= len(students); i++ {
	// 	if students[i] == "Ahmed" {
	// 		return i
	// 	}
	// }
	// return 0
}
