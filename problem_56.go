package main

import (
	"github.com/ivanthescientist/projecteuler/bignum"
	"fmt"
)

func main() {
	fmt.Println("Project Euler #56: largest digit sum number")

	maxSum := 0

	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			result := digitSum(generateNumber(i, j))
			if result > maxSum {
				maxSum = result
			}
		}
	}

	fmt.Println(digitSum(generateNumber(99, 99)))
	fmt.Println(maxSum)
}

func generateNumber(base, power int) []int {
	result := bignum.ToBigNumber(base)
	multiplier := bignum.ToBigNumber(base)

	for i := 0; i < power; i++ {
		result = bignum.Product(result, multiplier)
	}

	return result
}

func digitSum(num []int) int {
	sum := 0
	for _, digit := range num {
		sum += digit
	}

	return sum
}
