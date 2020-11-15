package student

func ActiveBits(n int) uint {
	count := 0

	if n < 0 {
		n *= -1
	}
	
	for n != 0 {
		if n % 2 == 1 {
			count++
		}
		n /= 2
	}

	return uint(count)
}
