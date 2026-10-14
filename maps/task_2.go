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
	fmt.Scan(&key)
	value, ok := students[key]
	reader := bufio.NewReader(os.Stdin)

	if ok {
		fmt.Println("yes student exist and name is :", value)
		fmt.Printf("Please anter the new name of the student")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)
		students[key] = newName
		fmt.Println("Updated name of the student is:", newName)
	} else {
		fmt.Println("Student not exist")
	}

}
