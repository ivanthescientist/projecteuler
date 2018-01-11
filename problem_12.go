package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/primes"
)

var primeList = primes.GenerateEratosthenes(1000000)

func main() {
	fmt.Println("Project Euler #12: Highly divisable triangular number")
	maxFactors := 0
	maxFactorsNumber := 0

	for i := 1; maxFactors < 500; i++ {
		current := getTriangularNumber(i)
		currentFactors := getTriangularNumberFactors(i)

		if currentFactors > maxFactors {
			maxFactors = currentFactors
			maxFactorsNumber = current
			fmt.Printf("New Max Divisors #%d for %dth Number: %d\n", maxFactors, i, maxFactorsNumber)
		}
	}

	fmt.Println(maxFactors)
	fmt.Println(maxFactorsNumber)
}

func getTriangularNumber(order int) int {
	return order * (order + 1) / 2
}

func getTriangularNumberFactors(n int) int {
	a := n
	b := n + 1
	aFactors := getPrimeFactorsList(a)
	bFactors := getPrimeFactorsList(b)

	if len(aFactors) == 0 || len(bFactors) == 0 {
		return 1
	}

	if a % 2 == 0{
		aFactors[0].Coefficient--
	}

	if b % 2 == 0 {
		bFactors[0].Coefficient--
	}

	return numberOfFactors(aFactors) * numberOfFactors(bFactors)
}

func numberOfFactors(array []PrimeFactor) int {
	product := 1

	for _, number := range array {
		product *= number.Coefficient
	}

	return product
}

type PrimeFactor struct {
	Prime int
	Coefficient int
}

func getPrimeFactorsList(number int) []PrimeFactor {
	factors := make([]PrimeFactor, 0)
	for _, prime := range primeList {
		if number < prime {
			break
		}
		if number % prime == 0 {
			current := number
			factor := PrimeFactor{
				Prime: prime,
				Coefficient: 1,
			}

			for ; current % prime == 0; current /= prime {
				factor.Coefficient++
			}
			factors = append(factors, factor)
		}
	}

	return factors
}
