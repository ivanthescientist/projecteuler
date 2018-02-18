package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/primes"
	"github.com/ivanthescientist/projecteuler/bignum"
	"github.com/ivanthescientist/projecteuler/permutations"
)

func main() {
	fmt.Println("Project Euler #35: Number of Circular Primes under 1000000")

	primeList := primes.GenerateEratosthenes(100)
	primeMap := make(map[int]bool)

	for _, prime := range primeList {
		primeMap[prime] = false
	}
	fmt.Println(len(primeMap))

	for prime, isCircular := range primeMap {
		if !isCircular {
			flagIfCircular(prime, primeMap)
		}
	}

	numberOfCirculars := 0

	for prime, isCircular := range primeMap {
		if isCircular {
			fmt.Println(prime)
			numberOfCirculars++
		}
	}

	fmt.Printf("There are %d circular primes between 0 and 1000000\n", numberOfCirculars)
}

func flagIfCircular(number int, primeMap map[int]bool) {
	digits := ToDigitsArray(number)
	perms := permutations.GenerateAll(digits)

	for _, perm := range perms {
		if _, exists := primeMap[ToNumber(perm)]; !exists {
			return
		}
	}

	for _, perm := range perms {
		primeMap[ToNumber(perm)] = true
	}
}

func ToDigitsArray(number int) []int {
	result := make([]int, 0)
	for ; number >= 1; {
		digit := number % 10
		result = append([]int{digit}, result... )
		number /= 10
	}

	return result
}

func ToNumber(digits []int) int {
	reversedDigits := bignum.Reverse(digits)
	result := 0

	for i, power := 0, 1; i < len(reversedDigits); i, power = i + 1, power * 10 {
		result += reversedDigits[i] * power
	}

	return result
}

// TODO generate all permutations of a number
// TODO add array to number function
// TODO flag all resulting prime numbers
// TODO count all non-false numbers
