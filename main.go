package main

import (
	"fmt"
)

func main() {
	fbg01()
	fizzBuzz2 := fbg02()
	fizzBuzz3 := fbg03()
	printFizzy(3, fizzBuzz3, fizzBuzz2)
}

func printFizzy(number int, fizzBuzzIs string, fizzBuzzWanted string) {
	if fizzBuzzIs == fizzBuzzWanted {
		fmt.Printf("%d is fizzy\n", number)
	} else {
		fmt.Printf("%d is not so fizzy %s\n", number, fizzBuzzIs)
	}
}
