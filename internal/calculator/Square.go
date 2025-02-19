package calculator

func Square(input uint32, square uint16) uint32 {

	if square == 0 {
		return 1
	}

	result := input
	for x := 1; x < int(square); x++ {
		result *= input
	}
	return result

}
