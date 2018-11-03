package main

import (
	"fmt"
)

func main() {
	fbg01()
	fizzBuzz2 := fbg02()
	printFizzy(3, fbg03, fizzBuzz2)
}

func printFizzy(number int, fizzBuzz func() string, fizzBuzzWanted string) {
	fizzBuzzIs := fizzBuzz()
	if fizzBuzzIs == fizzBuzzWanted {
		fmt.Printf("%d is fizzy\n", number)
	} else {
		fmt.Printf("%d is not so fizzy %s\n", number, fizzBuzzIs)
	}
}
