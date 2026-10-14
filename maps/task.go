package main

import (
	"fmt"
)

func classStudent() {
	students := make(map[int]string)

	students[1] = "Muhammad Asad"
	students[2] = "Ahmad Ali"
	students[3] = "Anam Zahra"

	rollNum := 0

	fmt.Println("Please enter the roll number of the student")
	fmt.Scan(&rollNum)
	// fmt.Println("Name of the student is :", students[roll_num])
	name, ok := students[rollNum]
	if ok {
		fmt.Println("Name of the student is :", name)
	} else {
		fmt.Println("Name not found", ok)
	}
}
