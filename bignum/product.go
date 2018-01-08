package bignum

func ProductRound(numberA []int, digitB int) []int {
	numberA = Reverse(numberA)
	result := make([]int, 0)
	overflow := 0

	for i := 0; i < len(numberA); i++ {
		digitA := GetDigit(numberA, i)
		product := digitA*digitB + overflow
		overflow = 0
		if product > 9 {
			overflow = product / 10
			product %= 10
		}
		result = append(result, product)
	}

	if overflow != 0 {
		result = append(result, overflow)
	}

	return Reverse(result)
}

func Product(numberA []int, numberB []int) []int {
	result := make([]int, 0)
	numberB = Reverse(numberB)

	for i := 0; i < len(numberB); i++ {
		roundResult := ProductRound(numberA, numberB[i])
		roundResult = append(roundResult, make([]int, i)...)
		result = Sum(result, roundResult)
	}

	return result
}
