package permutations

import (
	"testing"
)

func TestBeginningWalk(t *testing.T) {
	permutation := []int{1, 2, 3, 4}

	checkCase(t, NextPermutation(permutation), []int{1, 2, 4, 3})
	checkCase(t, NextPermutation(permutation), []int{1, 3, 2, 4})
	checkCase(t, NextPermutation(permutation), []int{1, 3, 4, 2})
	checkCase(t, NextPermutation(permutation), []int{1, 4, 2, 3})
}

func checkCase(t *testing.T, current[]int, expected []int) {
	if !slicesAreEqual(expected, current) {
		t.Errorf("Expected: %v, Got: %v", expected, current)
	}
}
