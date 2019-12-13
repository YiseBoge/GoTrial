package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less that %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is less that %d\n", y, x)
	} else {
		fmt.Printf("%d and %d are equals\n", x, y)
	}

	size := "big"
	switch size {
	case "big":
		fmt.Println("Huge ass Nigga")
	case "medium":
		fmt.Println("Not Bad")
	case "small":
		fmt.Println("Wazap tiny thing.")
	default:
		fmt.Println("Get your story straight.")
	}
}
