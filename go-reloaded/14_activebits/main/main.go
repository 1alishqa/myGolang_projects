package main

import (
	"fmt"
	student ".."
)

func main() {
	nbits := student.ActiveBits(-2)
	fmt.Println(nbits)

	nbits2 := student.ActiveBits(-15)
	fmt.Println(nbits2)

	nbits3 := student.ActiveBits(17)
	fmt.Println(nbits3)

	nbits4 := student.ActiveBits(4)
	fmt.Println(nbits4)

	nbits5 := student.ActiveBits(11)
	fmt.Println(nbits5)

	nbits6 := student.ActiveBits(9)
	fmt.Println(nbits6)

	nbits7 := student.ActiveBits(12)
	fmt.Println(nbits7)
}
