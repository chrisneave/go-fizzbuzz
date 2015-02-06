package main

import (
	"strconv"
)

func FizzBuzz(input int) (result string) {
	result = ""

	if input%3 == 0 {
		result += "Fizz"
	}

	if input%5 == 0 {
		result += "Buzz"
	}

	if len(result) == 0 {
		result = strconv.Itoa(input)
	}

	return result
}
