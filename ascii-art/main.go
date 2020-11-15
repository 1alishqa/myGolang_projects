package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	s := args[0]	// arg1 - string

	Runes := []rune(s)
	for _, v := range Runes {
		if v < 32 || v > 126 {
			fmt.Println("NO CYRILLIC LETTERS!")
			return
		}
	}

	fs := "standard.txt"	// add .txt , to open fs-file
	if len(args) > 1 {
		fs = args[1] + ".txt"
	}

	f, err := os.Open(fs)	// opening txt file
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	symbols := make([][]string, 95)	// arr IN arr, 95 symbols in txt
	var array []string
	var x int			// counting for every char in txt
	var z int			// index for arr of "symbols[][]"

	for scanner.Scan() {

		if scanner.Text() == "" {

			array = nil
			x = 0

		} else {

			array = append(array, scanner.Text())	// add text to array
			x++

			if x == 8 {
				symbols[z] = array		// assign the symbs to [][]
				z++

			} else if z == 95 {
				break			// break, when all symbs is assigned
			}

		}
	}

	array = make([]string, 8)
	var index int		// index, to find a symbol from [][]
	res := strings.Split(s, "\\n")

	for _, j := range res {

		for _, v := range j {
			index = int(rune(v)) - 32
			array = plusLine(array, symbols[index])
		}	// + every line for letters from args
	
		Printing(array)
		array = make([]string, 8)
	}
}

func Printing(arr []string) {

	for i := range arr {
		fmt.Println(arr[i])
	}
}

func plusLine(arr, elem []string) []string {

	for i := range arr {
		arr[i] += elem[i]
	}
	return arr
}
