package number

func GetDivisors(num int) []int {
	result := make([]int, 0)

	for i := 1; i < num; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}

	return result
}
