package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/number"
)

func main() {
	fmt.Println("Project Euler #21")

	var numbers = make(map[int]int)

	for i := 1; i < 10000; i++ {
		divisors := number.GetDivisors(i)
		amicableNumber := arraySum(divisors)
		amicableNumberDivisors := number.GetDivisors(amicableNumber)
		reverseSum := arraySum(amicableNumberDivisors)

		if reverseSum == i && i != amicableNumber {
			numbers[i] = amicableNumber
		}
	}

	sum := 0
	for origin, amicNum := range numbers {
		sum += amicNum

		fmt.Printf("d(%d) = %d\n", amicNum, origin)
	}

	fmt.Println(sum)
}

func arraySum(array []int) int {
	result := 0

	for _, value := range array {
		result += value
	}

	return result
}
