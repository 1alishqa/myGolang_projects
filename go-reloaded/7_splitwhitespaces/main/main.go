package main

import (
	"fmt"
	student ".."
)

func main() {
	str := "		Hello how are you?		"
	str2 := "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	str3 := "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	str4 := "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(student.SplitWhiteSpaces(str))
	fmt.Println(student.SplitWhiteSpaces(str2))
	fmt.Println(student.SplitWhiteSpaces(str3))
	fmt.Println(student.SplitWhiteSpaces(str4))
}
