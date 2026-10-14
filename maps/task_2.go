package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func listOfStudents() {
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
	reader := bufio.NewReader(os.Stdin)

	if ok {
		fmt.Println("yes student exist and name is :", value)
		fmt.Printf("Please anter the new name of the student : ")
		newName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unable to read the entered name")
			return
		}

		newName = strings.TrimSpace(newName)
		if newName == "" {
			fmt.Println("Name can not be empty ")
			return
		}
		students[key] = newName
		fmt.Println("Updated name of the student is : ", newName)
	} else {
		fmt.Println("Student not exist")
	}

}
