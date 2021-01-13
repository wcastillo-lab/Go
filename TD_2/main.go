package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	type User struct {
		Login    string `json:"Username"`
		Password string 
	}
	u := User{
		Login:     "Paul",
		Password:  "password123",
	}
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

}

