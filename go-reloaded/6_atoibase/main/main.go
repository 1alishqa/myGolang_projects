package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.AtoiBase("125a", "0123456789"))		// 0
	fmt.Println(student.AtoiBase("1111101", "01"))
	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

	fmt.Println()

	fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(student.AtoiBase("0001", "01"))
	fmt.Println(student.AtoiBase("00", "01"))
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
}
