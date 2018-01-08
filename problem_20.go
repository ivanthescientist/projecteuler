package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/bignum"
)

func main() {
	fmt.Println("Project Euler #20")
	fact100 := factorial(100)
	sum := 0
	for _, digit := range fact100 {
		sum += digit
	}
	fmt.Println(sum)
}

func factorial(n int) []int {
	result := []int{1}
	current := []int{1}

	for i := 0; i < n; i++ {
		result = bignum.Product(result, current)
		current = bignum.Sum(current, []int{1})
	}

	return result
}
