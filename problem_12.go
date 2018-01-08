package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/primes"
)

func main() {
	//current := 0
	//maxDivisors := 0
	//maxDivisorsNumber := 0
	//for i := 1; i < 100000; i++ {
	//	current += 1
	//	divisors := getDivisorsNumber(current)
	//	if divisors > maxDivisors {
	//		maxDivisors = divisors
	//		maxDivisorsNumber = i
	//	}
	//}
	//
	//println(maxDivisors)
	//println(maxDivisorsNumber)

	//fmt.Println(generatePrimes(1000))

	current := 0
	primeList := primes.GenerateEratosthenes(10000000000)
	maxDivisors := 0
	maxDivisorsNumber := 0

	startSearch := 700000

	for i := 1; i < 1000000; i++ {
		current += i
		if i > startSearch {
			divisors := getDivisorsNumber(i, primeList)
			if divisors > maxDivisors {
				maxDivisors = divisors
				maxDivisorsNumber = current
				fmt.Printf("New Max Divisors #%d for %dth Number: %d\n", maxDivisors, i, maxDivisorsNumber)
			}
		}
	}

	fmt.Println(maxDivisors)
	fmt.Println(maxDivisorsNumber)
}

func getDivisorsNumber(num int, primes []int) int {
	if num <= 1 {
		return 1
	}
	divisors := 1

	for _, prime := range primes {
		if prime > num {
			break
		}
		if num%prime == 0 {
			divisions := 0
			result := num
			for {
				if result%prime != 0 {
					break
				}

				result = result / prime
				divisions++
			}

			divisors *= divisions + 1
		}
	}

	return divisors
}
