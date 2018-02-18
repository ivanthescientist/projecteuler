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

func ToNumber(digits []int) int {
	reversedDigits := Reverse(digits)
	result := 0

	for i, power := 0, 1; i < len(reversedDigits); i, power = i + 1, power * 10 {
		result += reversedDigits[i] * power
	}

	return result
}

func ToBigNumber(value int) []int {
	current := value
	digits := make([]int, 0)

	for ;current > 0; {
		digit := current % 10
		digits = append(digits, digit)
		current /= 10
	}

	return Reverse(digits)
}