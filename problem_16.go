package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/bignum"
)

func main() {
	fmt.Println("Project Euler #16")

	fmt.Println(bignum.Product([]int{2, 9, 9}, []int{1, 0, 0, 0}))

	digitSum := 0
	for _, digit := range getPowerOfTwo(1000) {
		digitSum += digit
	}
	fmt.Println(digitSum)
}

func getPowerOfTwo(power int) []int {
	if power == 0 {
		return []int{1}
	}
	value := []int{2}

	for i := 1; i < power; i++ {
		value = bignum.Product(value, []int{2})
	}

	return value
}
