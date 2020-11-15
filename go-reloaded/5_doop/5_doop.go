package student

import (
	"github.com/01-edu/z01"
	// "z01"
	// "fmt"
)

// const max = 9223372036854775807
// const min = -9223372036854775808

func Doop(args []string) {
	count := 0
	for range args {
		count++
	}
	if count != 4 {
		return
	}
	var a, b int64
	var vrflw, vrflw2 bool
	var ans string
	str1 := args[1]
	str2 := args[3]

	if str1[0] == '+' {				// 1 - +11
		str1 = str1[1:]
	}

	if str2[0] == '+' {				// 
		str2 = str2[1:]
	}
	
	if isNumeric(str1) == false || isNumeric(str2) == false {
		Print("0")
		return
	}
	
	// a = int64(Atoi(str1))
	// b = int64(Atoi(str2))
	a, vrflw = Atoi(str1)
	b, vrflw2 = Atoi(str2)

	if !vrflw || !vrflw2 {
		Print("0")
		return
	}
	operator := args[2]
	
	if str1[0] != '-' && a < 0 || str1[0] == '-' && a > 0 || str2[0] != '-' && b < 0 || str2[0] == '-' && b > 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if operator == "+" {
		
		if a > 0 && b > 0 && a+b < 0 || a < 0 && b < 0 && a+b > 0 {
			Print(Itoa(0))
			return
		}

		ans = Itoa(a + b)

	} else if operator == "-" {
		
		if a < 0 && b > 0 && a-b > 0 || a > 0 && b < 0 && a-b < 0 {
			Print(Itoa(0))
			return
		}

		ans = Itoa(a - b)

	} else if operator == "*" {
		ans = mult(a, b)

	} else if operator == "/" {
		
		if a < 0 && b < 0 && a/b < 0 {
			Print(Itoa(0))
			return
		}
		if b == 0 {
			ans = "No division by 0"
		} else {
			ans = Itoa(a / b)
		}

	} else if operator == "%" {
		if b == 0 {
			ans = "No remainder of division by 0"
		} else {
			ans = Itoa(a % b)
		}

	} else {
		
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	Print(ans)
	/*for _, v := range ans {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')*/
}

func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}


func Itoa(n int64) string {
	var s string
	sign := true

	if n == -9223372036854775808 {
		s = "-9223372036854775808"
		return s
	}

	if n < 0 {
		n = n * -1
		sign = false
	}

	for {
		mod := n % 10
		s = string(mod+'0') + s
		n = n / 10
		if n == 0 {
			break
		}
	}

	if sign == false {
		s = "-" + s
	}
	// fmt.Println(s)
	return s
}

func mult(a, b int64) string {
	res := a * b
	
	if res/a != b {
		return Itoa(0)
	}
	return Itoa(a * b)
}

func isNumeric(str string) bool {

	if str == "0" {
		return true
	}

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if !IsNbr(s[i]) {
			return false
		}
	}
	return true

}

func IsNbr(s rune) bool {
	if s >= '0' && s <= '9' || s == '-' {
		return true
	}
	return false

}

func StrLen(str string) int {
	var result int
	Runes := []rune(str)

	for i := range Runes {
		result = i + 1
	}

	return result
}

func Atoi(s string) (int64, bool) {
	if s == "" {
		return 0, true
	}

	if s == "0" {
		return 0, true
	}
	
	// if s == "-9223372036854775808" {
	// 	return int64(-9223372036854775808), true
	// }

	Runes := []rune(s)
	sign := 1		// + or -

	if Runes[0] == '+' {

		sign = 1
		Runes = Runes[1:]
		if Overflow(string(Runes), "9223372036854775807") {
			return 0, false
		}

	} else if Runes[0] == '-' {

		sign = -1
		Runes = Runes[1:]
		if Overflow(string(Runes), "9223372036854775808") {
			return 0, false
		}

	} else {

		if Overflow(string(Runes), "9223372036854775807") {
			return 0, false
		}
	}

	var result int
	for _, i := range Runes {
		if i < 48 || i > 57 {
			return 0, false
		}
		result = result * 10 + (int(i) - 48)
	}

	// fmt.Println(result*sign)
	return int64(result * sign), true
}

func Overflow(s, over string) bool {
	if StrLen(s) > StrLen(over) {
		return true
	} else if StrLen(s) == StrLen(over) {

		for i := range s {
			if s[i] > over[i] {
				return true
			}
		}
	}
	return false
}
