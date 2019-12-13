package main

import "fmt"

func main() {

	fmt.Println("While looops with for loops for some reason:")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Another way is:")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("FizBuzzing:")
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
