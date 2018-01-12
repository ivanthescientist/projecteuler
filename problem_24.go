package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/permutations"
)

func main() {
	fmt.Println("Project Euler #24")

	var perms = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(perms)

	for i := 0; i < 1000000 + 1; i++ {
		permutations.NextPermutation(perms)
	}

	fmt.Printf("1,000,000-th permutation is: %+v\n", perms)
}



