package Humanity

import (
	"fmt"
	"strconv"
)

type Human struct {
	Name string
	Age	int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	buff, _ := strconv.Atoi(csv[1])
	human := Human{
		Name:    csv[0],
		Age:     buff,
		Country: csv[2],
	}
	fmt.Printf("%+v\n", human)
	return nil, nil
}