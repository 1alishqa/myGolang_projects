package student

func AtoiBase(s string, base string) int {

	len_base := Strlen(base)
	len_s := Strlen(s)
	valid := Validation(s, base)

	if !valid || len_base < 2 {
		return 0
	}

	var result int
	var res [70]string
	x := 0
	var sum int
	var decimal int

	for i := range s {
		for j := range base {

			if s[i] == base[j] {
				res[x] = Itoa(j)		// re-convert of "base"
				x++
				break
			}
		}
	}

	len_s--
	for v := range res {
		decimal = Atoi(string(res[v])) 		// each elem
		sum += (RecursivePower(len_base, len_s) * decimal)		// sum of all elem (convert to dec)
		len_s--
	}

	result = sum
	return result	
}

func Validation(s, base string) bool {
	valid := true
	var unique int

	for _, i := range s {
		repeat := 0
		for _, v := range base {
			if i == v {				// if "125a"
				repeat++
			}
		}

		if repeat == 0 {
			return false
		}
	}

	for _, i := range base {
		if i == '-' || i == '+' {
			valid = false
			break
		}

		unique = 0
		for _, v := range base {
			if i == v {
				unique++
			}
		}
		if unique > 1 {
			valid = false
			break
		}
	}
	return valid
}

func RecursivePower(nb int, power int) int {

	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	}

	res := nb * RecursivePower(nb, power-1)

	return res
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var res string
	for n != 0 {
		res = string(n % 10 + 48) + res
		n /= 10
	}
	return res
}

func Atoi(s string) int {

	if s == "" {
		return 0
	}

	Runes := []rune(s)
	var res int = 0

	for _, i := range Runes {
		res *= 10

		if i == '0' {
			res += 0
		} else if i == '1' {
			res += 1
		} else if i == '2' {
			res += 2
		} else if i == '3' {
			res += 3
		} else if i == '4' {
			res += 4
		} else if i == '5' {
			res += 5
		} else if i == '6' {
			res += 6
		} else if i == '7' {
			res += 7
		} else if i == '8' {
			res += 8
		} else if i == '9' {
			res += 9
		} else {
			return 0
		}
	}
	return res
}

func Strlen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
