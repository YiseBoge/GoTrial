package main

import "fmt"

func main() {
	//var name = "Yise Boge"
	name := "Yise Boge"
	var age uint32 = 21
	const isCool = true
	size := 30.5

	fmt.Println("Name:", name, "\nAge:", age)
	fmt.Printf("Coolness is of type: %T", isCool)
	fmt.Printf("We also have: %T", size)
}
