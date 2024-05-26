package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var filePath string

func init() {
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
}

func main() {
	var jp map[string]interface{}
	bText, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(bText)
	err = json.Unmarshal(bText, &jp)
	if err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}
}
