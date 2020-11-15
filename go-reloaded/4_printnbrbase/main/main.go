package main

import (
	"github.com/01-edu/z01"
	// "z01"
	student ".."
)

func main() {
	student.PrintNbrBase(0, "0123456789") // 0
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(1, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-1, "choumi") // -h
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "5")
	z01.PrintRune('\n')

	z01.PrintRune('\n')
	z01.PrintRune('\n')

	student.PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-661165, "1")
	z01.PrintRune('\n')
	student.PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}
