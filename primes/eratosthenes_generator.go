package primes

func GenerateEratosthenes(max int) []int {
	primes := make([]int, 0)
	isComposite := make([]bool, max)
	isComposite[0] = true
	isComposite[1] = true

	for i := 2; i*i < max; i++ {
		if !isComposite[i] {
			for j := i * 2; j < max; j += i {
				isComposite[j] = true
			}
		}
	}

	for i := 0; i < max; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
