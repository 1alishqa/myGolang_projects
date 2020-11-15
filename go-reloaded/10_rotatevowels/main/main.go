package main

import (
	"os"
	student ".."
)

func main() {
	args := os.Args[1:]
	student.RotateVowels(args)
}