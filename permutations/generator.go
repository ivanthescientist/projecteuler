package permutations

func GenerateAll(values []int) [][]int {
	permutations := make([][]int, 0)
	permutationsNum := PermutationsNumber(len(values))
	current := make([]int, len(values))
	copy(current, values)

	for i := 0; i < permutationsNum; i++ {
		copyOfCurrent := make([]int, len(current))
		copy(copyOfCurrent, current)
		permutations = append(permutations, copyOfCurrent)
		NextPermutation(current)
	}

	return permutations
}
