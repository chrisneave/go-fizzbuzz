package main

import "testing"
import "strconv"

func TestReturnsInputIfNotDivisibleBy3Or5(t *testing.T) {
	ins := []int{1, 2, 4, 7, 8}
	for _, i := range ins {
		actual := FizzBuzz(i)
		expected := strconv.Itoa(i)
		if actual != expected {
			t.Errorf("FizzBuzz(%d) = %s, should have been %s", i, actual, expected)
		}
	}
}

func TestReturnsFizzIfDivisibleBy3ButNot5(t *testing.T) {
	ins := []int{3, 6, 9, 12, 18, 21}
	for _, i := range ins {
		actual := FizzBuzz(i)
		expected := "Fizz"
		if actual != expected {
			t.Errorf("FizzBuzz(%d) = %s, should have been %s", i, actual, expected)
		}
	}
}

func TestReturnsBuzzIfDivisibleBy5ButNot3(t *testing.T) {
	ins := []int{5, 10, 20, 25, 35, 40}
	for _, i := range ins {
		actual := FizzBuzz(i)
		expected := "Buzz"
		if actual != expected {
			t.Errorf("FizzBuzz(%d) = %s, should have been %s", i, actual, expected)
		}
	}
}

func TestReturnsFizzBuzzIfDivisibleBy5And3(t *testing.T) {
	ins := []int{15, 30, 45, 60, 75, 90}
	for _, i := range ins {
		actual := FizzBuzz(i)
		expected := "FizzBuzz"
		if actual != expected {
			t.Errorf("FizzBuzz(%d) = %s, should have been %s", i, actual, expected)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(i)
	}
}
