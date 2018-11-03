package main

import (
	"fmt"
	"strconv"
)

// Standard
func fbg01() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(x)
		}
	}
}

// Build a string, minimize %
func fbg02() string {
	result := ""
	for x := 1; x <= 100; x++ {
		isFizz := x%3 == 0
		isBuzz := x%5 == 0
		if isFizz && isBuzz {
			result += "FizzBuzz"
		} else if isFizz {
			result += "Fizz"
		} else if isBuzz {
			result += "Buzz"
		} else {
			result += strconv.Itoa(x)
		}
		result += "\n"
	}
	return result
}

// Use range and case
func fbg03() string {
	result := ""
	for x := range [100]int{} {
		y := x + 1
		isFizz := y%3 == 0
		isBuzz := y%5 == 0
		switch isBuzz {
		case true:
			switch isFizz {
			case true:
				result += "FizzBuzz"
				break
			case false:
				result += "Buzz"
				break
			}
			break
		case false:
			switch isFizz {
			case true:
				result += "Fizz"
				break
			case false:
				result += strconv.Itoa(y)
				break
			}
			break
		}
		result += "\n"
	}
	return result
}
