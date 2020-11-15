/*
	Another fonts available on:
		github.com/drewnoakes/figgle/tree/master/fonts

	Some special characters DON't available on 3D font:
		(#, $, &, *, +, @, ~)

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
var black = "\033[38;2;0;0;0m%s\033[0m"
var red = "\033[38;2;255;0;0m%s\033[0m"
var green = "\033[38;2;0;128;0m%s\033[0m"
var yellow = "\033[38;2;255;255;0m%s\033[0m"
var blue = "\033[38;2;0;0;255m%s\033[0m"
var magenta = "\033[38;2;255;0;255m%s\033[0m"
var purple = "\033[38;2;128;0;128m%s\033[0m"
var cyan = "\033[38;2;0;255;255m%s\033[0m"
var gray = "\033[38;2;128;128;128m%s\033[0m"
var white = "\033[38;2;255;255;255m%s\033[0m"
var orange = "\033[38;2;255;140;0m%s\033[0m"*/

var threeD = false

func main() {
	args := os.Args[1:]
	s := args[0] // arg1 - string(text)
	if args[1] == "fontStyle/3d.txt" {
		threeD = true
	}
	// color := ReturnColor(args[2])

	f, _ := os.Open(args[1]) // opening txt file(fs)
	defer f.Close()

	symbols := ASCII(f)
	array := make([]string, 8)
	if threeD {
		array = make([]string, 16) //	height of every symbol = 16 in 3d
	}
	var index int // index, to find a symbol from [][]
	res := strings.Split(s, "\r\n")

	for _, j := range res {
		for _, v := range j {
			if (v < 32 || v > 126) && (v != 10) && (v != 13) {
				return
			}
			index = int(rune(v)) - 32
			array = plusLine(array, symbols[index])
		} // + every line for letters from args

		// Printing(array, color)
		Printing(array)
		array = make([]string, 8)
		if threeD {
			array = make([]string, 16)
		}
	}
}

func ASCII(f *os.File) [][]string {
	scanner := bufio.NewScanner(f)
	symbols := make([][]string, 95) // arr IN arr, 95 symbols in txt
	var array []string
	var x int // counting for every char in txt
	var z int // index for arr of "symbols[][]"

	for scanner.Scan() {
		if scanner.Text() == "" {
			array = nil
			x = 0

		} else {
			array = append(array, scanner.Text()) // add text to array
			x++
			if threeD {
				if x == 16 {
					symbols[z] = array // assign the symbs to [][]
					z++
				} else if z == 95 {
					break // break, when all symbs is assigned
				}

			} else {
				if x == 8 {
					symbols[z] = array // assign the symbs to [][]
					z++
				} else if z == 95 {
					break // break, when all symbs is assigned
				}
			}
		}
	}
	return symbols
}

func Printing(arr []string /*, color string*/) {

	for i := range arr {
		fmt.Println(arr[i])
		// fmt.Printf(color, arr[i])
		// fmt.Println()
	}
}

func plusLine(arr, elem []string) []string {

	for i := range arr {
		arr[i] += elem[i]
	}
	return arr
}

/*
func ReturnColor(color string) string {
	var letterColor string
	if color == "orange" {
		letterColor = orange
	} else if color == "red" {
		letterColor = red
	} else if color == "black" {
		letterColor = black
	} else if color == "green" {
		letterColor = green
	} else if color == "yellow" {
		letterColor = yellow
	} else if color == "blue" {
		letterColor = blue
	} else if color == "magenta" {
		letterColor = magenta
	} else if color == "purple" {
		letterColor = purple
	} else if color == "cyan" {
		letterColor = cyan
	} else if color == "gray" {
		letterColor = gray
	} else if color == "white" {
		letterColor = white
	}

	return letterColor
}*/
