package student

import (
	"github.com/01-edu/z01"
	// "z01"
	// "os"
)

func RotateVowels(args []string) {
	length := 0
	for range args {
		length++
	}

	if length != 0 {

		var str string		// str - all args
		var vowel string	// vowels in args
		var len_vwl int
		var not_1st bool = false	// not first

		for _, s := range args {
			if not_1st {
				str += " "
			}
			not_1st = true
			str += s		// add all args
		}

		Runes := []rune(str)

		for i := range Runes {

			if isVowel(Runes[i]) {
				vowel += string(Runes[i])	// all vowels in args
				len_vwl++
			}
		}

		for i := range Runes {

			if isVowel(Runes[i]) {
				len_vwl--
				Runes[i] = rune(vowel[len_vwl])		// subtition
			}
		}

		for _, i := range Runes {
			z01.PrintRune(i)
		}

	}
	z01.PrintRune(10)
}

func isVowel(s rune) bool {
	if (s == 'a' || s == 'A' || s == 'e' || s == 'E' || 
		s == 'u' || s == 'U' || s == 'i' || s == 'I' || s == 'o' || s == 'O') {

			return true
		}

	return false
}
