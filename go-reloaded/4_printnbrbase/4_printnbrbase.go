package student

import (
	"github.com/01-edu/z01"
	// "z01"
)

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func PrintNbrBase(nbr int, base string) {

	var res string
	len_base := Strlen(base)
	valid := Validation(base)
	overflow := false
	
	if !valid || len_base < 2 {
		for _, i := range "NV" {
			z01.PrintRune(i)
		}
		return
	}

	if MaxInt32 < nbr || MinInt32 > nbr {		// overflows
		overflow = true
	}
	/*
	if nbr == -0 {
		z01.PrintRune('-')
		z01.PrintRune(rune(base[0]))
	}*/

	if nbr == 0 {		// template
		z01.PrintRune(rune(base[0]))
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}

	if overflow {
		Over(nbr, base, len_base)
		return
	}
	
	for nbr > 0 {
		res += string(base[nbr % len_base])
		nbr /= len_base
	}

	res = Reverse(res)
	for _, i := range res {
		z01.PrintRune(i)
	}
}

func Reverse(s string) string {
	len := Strlen(s)
	var res string

	for i := len - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func Validation(base string) bool {
	valid := true
	var unique int
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

func Strlen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func Over(nbr int, base string, len_base int) {

	len_nbr := 0
	temp := nbr
	for temp != 0 {
		len_nbr++
		temp /= len_base
	}

	var Nums [70]int
	var Runes [70]rune

	Nums[0] = 9
	Nums[1] = 2
	Nums[2] = 2
	Nums[3] = 3
	Nums[4] = 3
	Nums[5] = 7
	Nums[6] = 2
	Nums[7] = 0
	Nums[8] = 3
	Nums[9] = 6
	Nums[10] = 8
	Nums[11] = 5
	Nums[12] = 4
	Nums[13] = 7
	Nums[14] = 7
	Nums[15] = 5
	Nums[16] = 8
	Nums[17] = 0
	Nums[18] = 8

	
	base_runes := []rune(base)
	x := 0

	for i := 0; i < len_nbr; i++ {
		for j := range base_runes {

			if j == Nums[i] {
				Runes[x] = base_runes[j]
				x++
			}

		}
	}

	for i := 0; i < len_nbr; i++ {
		z01.PrintRune(Runes[i])
	}
}
