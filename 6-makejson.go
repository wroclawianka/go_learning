package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var userInfo = make(map[string]string)
	var name, address string

	fmt.Println("Enter your name : ")
	fmt.Scan(&name)
	userInfo["name"] = name

	fmt.Println("Enter your address : ")
	fmt.Scan(&address)
	userInfo["address"] = address

	if jsonData, err := json.Marshal(userInfo); err == nil {
		fmt.Println(string(jsonData))
	}
}
