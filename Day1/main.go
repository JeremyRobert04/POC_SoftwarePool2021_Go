package main

import (
	"SofwareGoDay1/Humanity"
	"SofwareGoDay1/data"
	"fmt"
)

func main() {
	content, err := data.ReadFile("list")
	if err != nil {
		fmt.Println(err)
	}
	for i := range content {
		result, _ := data.LineToCSV(content[i])
		Humanity.NewHumanFromCSV(result)
	}
}