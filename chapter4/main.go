package main

import "fmt"

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed
}

func main() {
	monkey := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 3,
		},
	}
	fmt.Printf("%v\n", monkey.Amount)
}
