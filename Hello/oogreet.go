package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Hi " + p.name + " !"
}

func main() {
	greeter := Person{"Ece"}
	fmt.Println(greeter.greet())
}
