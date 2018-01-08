package main

import "fmt"

func main() {
	fmt.Println("Project Euler #14")

	longest := 0
	longestNumber := 0

	for i := 2; i < 1000000; i++ {
		seqLen := countSequence(i)

		if seqLen > longest {
			longest = seqLen
			longestNumber = i
		}
	}

	fmt.Printf("Number %d with sequence of %d", longestNumber, longest)
}

func countSequence(start int) int {
	if start < 2 {
		return 0
	}

	result := start
	steps := 1

	for result != 1 {
		if result%2 == 0 {
			result /= 2
		} else {
			result = result*3 + 1
		}
		steps++
	}

	return steps
}
