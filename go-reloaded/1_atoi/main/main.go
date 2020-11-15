package main

import (
	"fmt"
	student ".."
)

func main() {
	s := "123a"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello5World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "mf"
	s9 := "-9223372036854775807"
	s10 := "9223372036854775808"

	n := student.Atoi(s)
	n2 := student.Atoi(s2)
	n3 := student.Atoi(s3)
	n4 := student.Atoi(s4)
	n5 := student.Atoi(s5)
	n6 := student.Atoi(s6)
	n7 := student.Atoi(s7)
	n8 := student.Atoi(s8)
	n9 := student.Atoi(s9)
	n10 := student.Atoi(s10)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)
}