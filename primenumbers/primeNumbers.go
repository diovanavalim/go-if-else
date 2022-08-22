package primenumbers

import "math"

func IsPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; float64(i) <= math.Sqrt(float64(number)); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
