package main

import (
	"fmt"
	// "github.com/01-edu/z01"
	// "z01"
	"os"
)


func main() {
	args := os.Args[1:]
	var commands []string
	var errors []string
	next := false
	change := false

	if lenArr(args) >= 1 {

		for _, i := range args {

			for _, j := range SplitWhiteSpaces(i) {

				if next {
					if change {
						commands[1] = j

					} else {

						if lenArr(commands) > 1 {
							commands = append(commands, commands[1])
							commands[1] = j

						} else {
							commands = append(commands, j)
						}
					}

					next = false
					continue
				}

				if j == "-c" {
					next = true
					if lenArr(commands) > 0 {

						if commands[0] == "-c" {
							change = true
							continue
						}

						commands = append(commands, commands[0])
						commands[0] = j

					} else {
						commands = append(commands, j)
					}

				} else if lenStr(j) >= 2 && j[:2] == "-c" {

					if lenArr(commands) > 0 {
						if commands[0] == "-c" {
							commands[0] = j
							continue
						}
						commands = append(commands, commands[0])
						commands[0] = j

					} else {
						commands = append(commands, j)
					}

				} else {
					commands = append(commands, j)
				}

			}
		}

		if commands[0] != "-c" && commands[0][:2] != "-c" {
			os.Exit(1)
		}


		var num int
		var err bool

		if lenStr(commands[0]) > 2 {
			num, err = Atoi(commands[0][2:])
			if err {
				fmt.Printf("%v", "tail: invalid number of bytes: '"+commands[0][2:]+"'\n")
				os.Exit(1)
			}
			commands = commands[1:]

		} else {

			num, err = Atoi(commands[1])
			if err {
				fmt.Printf("%v", "tail: invalid number of bytes: '"+commands[1]+"'\n")
				os.Exit(1)
			}

			commands = commands[2:]
		}

		printName := lenArr(commands) > 1

		for j, f := range commands {

			fi, err := os.Open(f)
			if err != nil {
				splitErr := Split(err.Error(), ": ")
				formatedErr := Capitalize(splitErr[lenArr(splitErr)-1:][0])
				errors = append(errors, "tail: cannot open '"+f+"' for reading: "+formatedErr+"\n")
				continue
			}

			statErr, size := fileSize(fi)
			if statErr {
				continue
			}

			if printName {
				fmt.Printf("==> %s <==\n", f)
			}

			var read []byte
			if num > int(size) {
				num = int(size)
			}

			for i := 0; i < num; i++ {
				read = append(read, 0x0)
			}
			_, er := fi.ReadAt(read, size-int64(num))
			if er != nil {
				errors = append(errors, err.Error())
				continue
			}

			for _, c := range read {
				fmt.Printf("%c", rune(c))
			}

			if j < lenArr(commands)-1 {
				fmt.Printf("\n")
			}
			fi.Close()
		}
		if lenArr(errors) > 0 {
			for _, v := range errors {
				fmt.Printf("%v", v)
			}
			os.Exit(1)
		}
	}
}

func Capitalize(s string) string {

	nextCap := true
	arr := []rune(s)
	ans := ""
	for _, v := range arr {

		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {

			if nextCap {

				if v >= '0' && v <= '9' {
					ans += string(v)

				} else {

					if v >= 'a' && v <= 'z' {
						ans += string(v - 32)
					} else {
						ans += string(v)
					}
				}

				nextCap = false

			} else {

				if v >= 'A' && v <= 'Z' {
					ans += string(v + 32)
				} else {
					ans += string(v)
				}
			}

		} else {
			nextCap = true
			ans += string(v)
		}
	}
	return ans
}

func lenArr(arr []string) int {
	cnt := 0
	for range arr {
		cnt++
	}
	return cnt
}

func lenStr(str string) int {
	cnt := 0
	for range str {
		cnt++
	}
	return cnt
}

func Split(str, charset string) []string {

	i := 0
	word := ""
	arrLen := 1
	cnt := 0
	strLen := lenStr(str)
	charsetLen := lenStr(charset)

	for i != strLen-1 {

		if strLen-1 > i+charsetLen && str[i:i+charsetLen] == charset {
			arrLen++
		}
		i++
	}

	ans := make([]string, arrLen)
	i = 0
	for i != strLen {

		if strLen-1 > i+charsetLen && str[i:i+charsetLen] == charset {

			if lenStr(word) > 0 {
				ans[cnt] = word
				word = ""
				cnt++
			}
			i += charsetLen

		} else {
			word += string(str[i])
			i++
		}
	}
	ans[arrLen-1] = word
	return ans
}

func Overflow(a int, b int, c int) bool {

	result := a*b + c
	if a > 0 && b > 0 && result < 0 {
		return true
	}
	if a < 0 && b < 0 && result > 0 {
		return true
	}
	return false

}

func Atoi(s string) (int, bool) {
	
	p := []rune(s)
	var ans int
	sign := true
	for index, i := range p {
		if index == 0 && (i == '-' || i == '+') {
			if i == '-' {
				sign = false
			} else {
				sign = true
			}
		} else if index != 0 && (i == '-' || i == '+') {
			return 0, true
		} else {
			if int(i-'0') > 9 || int(i-'0') < 0 {
				return 0, true
			}
			if Overflow(ans, 10, int(i-'0')) {
				return 0, true
			}
			ans = ans*10 + int(i-'0')
		}

	}
	_ = sign
	return ans, false
}

func makeArr(n int) []string {
	var arr []string
	for i := 0; i < n; i++ {
		arr = append(arr, "")
	}
	return arr
}

func SplitWhiteSpaces(str string) []string {
	str = str + " "
	space := false
	arrayLen := 0
	for _, i := range str {
		if i == '\n' || i == '\t' {
			arrayLen++
		}
		if i == ' ' && !(space) {
			space = true
			arrayLen++
		} else if i != ' ' && i != '\n' && i != '\t' {
			space = false
		}
	}
	tmp := ""
	cnt := 0
	arr := makeArr(arrayLen)
	counter := 0
	for _, i := range str {
		if i != '\n' && i != '\t' && i != ' ' {
			tmp += string(i)
			counter++
		} else {
			if counter != 0 {
				counter = 0
				arr[cnt] = tmp
				cnt++
				tmp = ""
			}
		}
	}
	return arr
}

func fileSize(fi *os.File) (bool, int64) {
	fil, err := fi.Stat()

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return true, 0
	}

	return false, fil.Size()
}
