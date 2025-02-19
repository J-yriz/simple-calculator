package calculator

func Factorial(input uint32) uint32 {

	if input == 0 || input == 1 {
		return 1
	}

	return input * Factorial(input-1)

}
