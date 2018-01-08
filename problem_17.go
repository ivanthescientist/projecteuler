package main

import "fmt"

var first = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var tens = []string{
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var second = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func main() {
	fmt.Println(extractDigits(105))
	fmt.Println(numberToText(1))
	fmt.Println(numberToText(2))
	fmt.Println(numberToText(9))
	fmt.Println(numberToText(11))
	fmt.Println(numberToText(19))
}

func numberToText(number int) string {
	result := ""
	digits := extractDigits(number)

	// TODO: split into several functions each handling specific cases and then modularize code below out of them.
	if len(digits) == 1 {
		return first[digits[0]]
	} else if len(digits) == 2 {
		if digits[0] == 1 {
			return tens[digits[1]]
		} else {

		}

	} else if len(digits) == 3 {

	} else {
		return "one thousand"
	}

	return result
}

func extractDigits(number int) []int {
	result := make([]int, 0)

	for number >= 1 {
		result = append([]int{number % 10}, result...)
		number = number / 10
	}

	return result
}
