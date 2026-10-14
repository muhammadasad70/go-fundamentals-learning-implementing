package main

import (
	"fmt"
)

func deleteStudent() {
	students := map[int]string{
		1: "Muhammad Ahmad",
		2: "Aoun Abbas",
		3: "Maaveer ",
	}
	key := 0
	fmt.Println("Please enter the key of the student ")
	_, err := fmt.Scan(&key)
	if err != nil {
		fmt.Println("Invalid student id")
		return
	}
	value, ok := students[key]
	if ok {
		fmt.Println("Found student and name is :", value)
		delete(students, key)
		fmt.Println("Student wth key has been deleted successfully ", key)
		if _, exist := students[key]; !exist {
			fmt.Println("Student deletion verified successfully")
		}
		fmt.Println("Remaining students are:")
		for id, name := range students {
			fmt.Println(id, " ", name)
		}

	} else {
		fmt.Println("Unable to find the student")
	}

}
