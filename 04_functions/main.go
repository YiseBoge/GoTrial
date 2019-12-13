package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(greet("Abebe"))
	fmt.Println(add(1, 2))
}
