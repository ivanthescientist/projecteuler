package permutations

import "github.com/ivanthescientist/projecteuler/util"

func PermutationsNumber(length int) int {
	if length == 1 {
		return 1
	}
	return length * PermutationsNumber(length-1)
}

func NextPermutation(values []int) []int {
	lastIndex := len(values) - 1
	i := lastIndex - 1

	for i >= 0 && values[i] > values[i+1] {
		i--
	}

	if i < 0 {
		util.ReverseSublist(values, 0, lastIndex)
	} else {
		j := lastIndex
		for j > 0 && values[j] < values[i] {
			j--
		}
		values[i], values[j] = values[j], values[i]
		util.ReverseSublist(values, i+1, lastIndex)
	}
	return values
}
