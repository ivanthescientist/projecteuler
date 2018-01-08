package bignum

import "strconv"

func ToString(number []int) string {
	result := ""
	for _, digit := range number {
		result += strconv.Itoa(digit)
	}
	return result
}

func Reverse(array []int) []int {
	result := make([]int, len(array))
	copy(result, array)
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func GetDigit(number []int, index int) int {
	if index >= len(number) {
		return 0
	}

	return number[index]
}

func GetMaxLength(numberA []int, numberB []int) int {
	lenA := len(numberA)
	lenB := len(numberB)

	if lenB > lenA {
		return lenB
	} else {
		return lenA
	}
}
