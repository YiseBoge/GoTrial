package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", ""} // slice when it doesnt have fixed size

	fmt.Println(fruitArr)
	fmt.Println(planets)
	fmt.Println("Planets Number is:", len(planets))
	fmt.Println("Inner Planets are:", planets[0:4])
	fmt.Println("Hey!! we arrays can slice tooooo:", fruitArr[0:1])
}
