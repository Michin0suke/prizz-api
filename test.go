package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	name string
	age int
}

func main() {
	resp := `
	{
		name: "michiya",
		age: 10
	}
	`
	var u user
	json.Unmarshal(byte[](resp), &u)
	fmt.Println(u)
}
