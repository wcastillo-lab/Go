package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	var jsonUser = []byte(
	
	`[
		{"userName": "matm", "Password": "123456"}, 
		{"userName": "fake44", "Password": "azerty"}	
	]`)	
	var users []User
	err := json.Unmarshal(jsonUser, &users)
	if err != nil {
		fmt.Println("error:", err)

	}
	fmt.Printf("%+v\n", users)

}
type User struct {
        Login    string `json:"Username"`
        Password string
        }

