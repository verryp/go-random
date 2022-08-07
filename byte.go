package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	test := "{}"
	testByte, _ := json.Marshal(test)

	var user *User
	_ = json.Unmarshal(testByte, &user)

	fmt.Println(len(testByte))
	fmt.Println(user.Age)
	fmt.Println(user == nil)
}
