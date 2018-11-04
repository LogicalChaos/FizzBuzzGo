package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

func main() {
	fbg01()
	fizzBuzz2 := fbg02()
	printFizzy(3, fbg03, fizzBuzz2)
	printFizzy(4, fbg04, fizzBuzz2)
}

func printFizzy(number int, fizzBuzz func() string, fizzBuzzWanted string) {
	fizzBuzzIs := fizzBuzz()
	if fizzBuzzIs == fizzBuzzWanted {
		fmt.Printf("%d is fizzy\n", number)
	} else {
		diffs := dmp.DiffMain(fizzBuzzWanted, fizzBuzzIs, false)
		fmt.Printf("%d is not so fizzy %s\n", number, diffs)
	}
}
