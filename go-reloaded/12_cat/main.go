package main

import (
	// "github.com/01-edu/z01"
	"z01"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0
	for range args {
		count++
	}

	if count == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return
		}
	}

	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			for _, i := range err.Error() {
				z01.PrintRune(i)
			}
			z01.PrintRune(10)
			return
		}

		another, err := ioutil.ReadAll(file)
		for _, i := range string(another) {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)

		if err != nil {
			for _, i := range err.Error() {
				z01.PrintRune(i)
			}
			z01.PrintRune(10)
		}
	}
}
