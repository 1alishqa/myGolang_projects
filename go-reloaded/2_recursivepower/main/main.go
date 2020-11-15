package main

import (
	"fmt"
	student ".."
)

func main() {
	arg1 := 0
	arg2 := 0
	fmt.Println(student.RecursivePower(arg1, arg2))

	arg3 := -1
	arg4 := 1
	fmt.Println(student.RecursivePower(arg3, arg4))

	arg5 := -6
	arg6 := 5
	fmt.Println(student.RecursivePower(arg5, arg6))
}