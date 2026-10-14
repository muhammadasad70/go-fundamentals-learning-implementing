package main

import (
	"fmt"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

func studentRecord() {
	students := make(map[int]Student)

	// students:=map[int]student{
	// 	1:{
	// 		ID: 1,
	// 		Name:"Muhammad Asad",
	// 		Age: 24,
	// 	},
	// 	2:{
	// 		ID: 2,
	// 		Name:"Aoun Abbas",
	// 		Age: 23,
	// 	},
	// 	3:{
	// 		ID: 3,
	// 		Name:"Hassan Nawaz",
	// 		Age: 22,
	// 	},
	// }

	students[1] = Student{
		ID:   1,
		Name: "Muhammad Asad",
		Age:  24,
	}
	students[2] = Student{
		ID:   2,
		Name: "Aoun Abbas",
		Age:  23,
	}
	students[3] = Student{
		ID:   3,
		Name: "Hassan Nawaz",
		Age:  22,
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
		fmt.Println("Id : ", value.ID)
		fmt.Println("Name of the student : ", value.Name)
		fmt.Println("Age of the student : ", value.Age)
	} else {
		fmt.Println("Student not found")
	}

}
