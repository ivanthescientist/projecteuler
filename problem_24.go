package main

import "fmt"

func main() {
	fmt.Println("Project Euler #24")

	var perms = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(perms)

	for i := 0; i < 1000000+1; i++ {
		fmt.Printf("Perm #%d: %+v\n", i+1, perms)
		nextPermutation(perms)
	}
}

func permutationsNumber(length int) int {
	if length == 1 {
		return 1
	}
	return length * permutationsNumber(length-1)
}

func nextPermutation(values []int) []int {
	lastIndex := len(values) - 1
	i := lastIndex - 1

	for i >= 0 && values[i] > values[i+1] {
		i--
	}

	if i < 0 {
		reverseSublist(values, 0, lastIndex)
	} else {
		j := lastIndex
		for j > 0 && values[j] < values[i] {
			j--
		}
		values[i], values[j] = values[j], values[i]
		reverseSublist(values, i+1, lastIndex)
	}
	return values
}

func reverseSublist(list []int, start int, end int) {
	if start < end {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	}
}
