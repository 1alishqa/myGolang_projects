package student

import (
	// "z01"
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	} 

	count := 0

	for i := 0; i < 10; i++ {

		if n > 1 {
			Comb(n-1, i, Itoa(i), &count)

		} else {					// for n = 1
			if count > 0 {
				Print(", ")
			}
			Print(Itoa(i))
			count++
		}
	}
	Print("\n")
}

func Comb(nbr int, prev int, res string, count *int) {	// result for 1st digit or nbr
	
	for i := 0; i < 10; i++ {									// prev = shouldn't start < than prev
		if prev < i {

			if nbr == 1 {
				if *count > 0 {
					Print(", ")
				}

				Print(res + Itoa(i))
				*count++

			} else {
				Comb(nbr - 1, i, res + Itoa(i), count)
			}
		}
	}
}

func Itoa(nbr int) string {
	s := ""

	if nbr == 0 {
		return "0"
	}

	for nbr > 0 {
		s = string(nbr % 10 + 48) + s
		nbr /= 10
	}
	return s
}

func Print(s string) {

	for _, i := range s {
		z01.PrintRune(i)
	}
}
