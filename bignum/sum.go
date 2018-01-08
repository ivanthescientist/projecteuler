package bignum

func Sum(numberA []int, numberB []int) []int {
	numberA = Reverse(numberA)
	numberB = Reverse(numberB)
	maxLen := GetMaxLength(numberA, numberB)
	result := make([]int, 0)
	overflow := 0

	for i := 0; i < maxLen; i++ {
		digitA := GetDigit(numberA, i)
		digitB := GetDigit(numberB, i)

		sum := digitA + digitB + overflow
		overflow = 0
		if sum > 9 {
			overflow = 1
			sum %= 10
		}
		result = append(result, sum)
	}

	if overflow != 0 {
		result = append(result, 1)
	}

	return Reverse(result)
}
