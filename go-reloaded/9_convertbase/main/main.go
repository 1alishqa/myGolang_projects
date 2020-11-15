package main

import (
	"fmt"
	student ".."
)

func main() {
	result := student.ConvertBase("10f1011", "01", "0123456789")	// NV
	fmt.Println(result)
	result2 := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	fmt.Println(result2)
	result3 := student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	fmt.Println(result3)
	result4 := student.ConvertBase("256850", "0123456789", "01")
	fmt.Println(result4)
	result5 := student.ConvertBase("uuhoumo", "choumi", "Zone01")
	fmt.Println(result5)
	result6 := student.ConvertBase("9223372036854775807", "0123456789", "0123456789")
	fmt.Println(result6)
}
