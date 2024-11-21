package main

import "testing"

func TestFizz(t *testing.T) {
	table := []struct {
		number int
		result string
	}{
		{number: 1, result: "1"},
		{number: 3, result: "Fizz"},
		{number: 5, result: "Buzz"},
		{number: 10, result: "Buzz"},
		{number: 15, result: "FizzBuzz"},
		{number: 17, result: "17"},
		{number: 18, result: "Fizz"},
		{number: 25, result: "Buzz"},
		{number: 30, result: "FizzBuzz"},
		{number: 45, result: "FizzBuzz"},
		{number: 67, result: "67"},
		{number: 75, result: "FizzBuzz"},
		{number: 88, result: "88"},
		{number: 100, result: "Buzz"},
	}

	for _, v := range table {
		result := FizzBuzz(v.number)
		if result != v.result {
			t.Errorf("%d is %s and should be %s", v.number, result, v.result)
		}
	}
}
