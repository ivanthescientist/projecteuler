package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/primes"
	"github.com/ivanthescientist/projecteuler/permutations"
	"github.com/ivanthescientist/projecteuler/bignum"
)

func main() {
	fmt.Println("Project Euler #41: largest pandigital prime")

	var largestPrimePandigital = 0
	var pandigitals = [][]int{
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 7, 8,},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	var primeList = primes.GenerateEratosthenes(123456789)
	var primeSet = make(map[int] struct{})
	for _, prime := range primeList {
		primeSet[prime] = struct{}{}
	}

	for _, pandigital := range pandigitals {
		perms := permutations.GenerateAll(pandigital)
		for _, perm := range perms {
			num := bignum.ToNumber(perm)
			if _, isPresent := primeSet[num]; isPresent && num > largestPrimePandigital {
				largestPrimePandigital = num
			}
		}
	}

	fmt.Println(len(primeList))
	fmt.Printf("%d", largestPrimePandigital)
}