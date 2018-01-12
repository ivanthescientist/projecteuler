package permutations

import (
	"testing"
)

func TestGenerationFromLXGStart(t *testing.T) {
	expectedPermutations := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	generatedPermutations := GenerateAll([]int{1, 2, 3})

	for i := 0; i < len(expectedPermutations); i++ {
		if !slicesAreEqual(expectedPermutations[i], generatedPermutations[i]) {
			t.Errorf("Wrong %d Slice: %+v != %+v", i, expectedPermutations[i], generatedPermutations[i])
		}
	}
}

func TestGenerationFromLXGMiddle(t *testing.T) {
	expectedPermutations := [][]int{
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
		{1, 2, 3},
		{1, 3, 2},
	}
	generatedPermutations := GenerateAll([]int{2, 1, 3})

	for i := 0; i < len(expectedPermutations); i++ {
		if !slicesAreEqual(expectedPermutations[i], generatedPermutations[i]) {
			t.Errorf("Wrong %d Slice: %+v != %+v", i, expectedPermutations[i], generatedPermutations[i])
		}
	}
}

func slicesAreEqual(sliceA []int, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
