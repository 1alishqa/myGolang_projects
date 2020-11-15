package student

func Atoi(s string) int {

	if s == "" {
		return 0
	}
	
	if s == "-9223372036854775808" {
		return -9223372036854775808
	}

	Runes := []rune(s)
	sign := 1		// + or -

	if Runes[0] == '+' {

		sign = 1
		Runes = Runes[1:]
		if Overflow(string(Runes), "9223372036854775807") {
			return 0
		}

	} else if Runes[0] == '-' {

		sign = -1
		Runes = Runes[1:]
		if Overflow(string(Runes), "9223372036854775808") {
			return 0
		}

	}

		if Overflow(string(Runes), "9223372036854775807") {
			return 0
		}

	var result int
	for _, i := range Runes {
		if i < 48 || i > 57 {
			return 0
		}
		result = result * 10 + (int(i) - 48)
	}

	return result * sign

}

func Overflow(s, over string) bool {
	if Strlen(s) > Strlen(over) {
		return true
	} else if Strlen(s) == Strlen(over) {

		for i := range s {
			if s[i] > over[i] {
				return true
			}
		}
	}
	return false
}

func Strlen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}