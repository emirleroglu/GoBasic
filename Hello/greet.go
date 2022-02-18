package main

import "fmt"

func main() {
	var name string = "Halil"
	var greeting = createGreet(name)
	fmt.Println(greeting)

}
func createGreet(name string) string {
	greeting := "Hi " + name + " !"
	return greeting
}
