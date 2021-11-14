package main

import (
	"chapter2/animals"
	"fmt"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println((animals.MonkeyFeed()))
	fmt.Println(animals.RabbitFeed())
}
