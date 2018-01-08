package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var alphabet = map[rune]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'D': 4,
	'E': 5,
	'F': 6,
	'G': 7,
	'H': 8,
	'I': 9,
	'J': 10,
	'K': 11,
	'L': 12,
	'M': 13,
	'N': 14,
	'O': 15,
	'P': 16,
	'Q': 17,
	'R': 18,
	'S': 19,
	'T': 20,
	'U': 21,
	'V': 22,
	'W': 23,
	'X': 24,
	'Y': 25,
	'Z': 26,
}

func main() {
	fmt.Println()

	names := parseInput(readFileInputFile())
	sort.Strings(names)

	scoreSum := 0

	for index, name := range names {
		scoreSum += alphaSum(name) * (index + 1)
	}

	fmt.Printf("Score sum is %d\n", scoreSum)
}

func alphaSum(value string) int {
	sum := 0

	for _, letter := range value {
		sum += alphabet[letter]
	}

	return sum
}

func parseInput(input string) []string {
	input = strings.Replace(input, "\"", "", -1)
	names := strings.Split(input, ",")

	return names
}

func readFileInputFile() string {
	fileBytes, _ := ioutil.ReadFile("./data/problem_22_data.txt")
	return string(fileBytes)
}
