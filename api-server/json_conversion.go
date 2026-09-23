package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonConversion() {
	user := User{
		Name: "Asad",
		Age:  24,
	}
	// encoding the go data into json bytes because we use Marshal()when we have to encode and decode data as a byte like we use encoding and decoding when we need data as a stream

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}
	fmt.Println("Converting the go data into json :")
	fmt.Println(string(data))
	// now we are decoing the json []byte data into the go data by using the Unmarshal()

	var decodeUser User

	err = json.Unmarshal(data, &decodeUser)
	if err != nil {
		fmt.Println("Unmarshal Error", err)
		return
	}
	fmt.Println("Converting the json data into go data :")
	fmt.Println(decodeUser.Name)
	fmt.Println(decodeUser.Age)
}
