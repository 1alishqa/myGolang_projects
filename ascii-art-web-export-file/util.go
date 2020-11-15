package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getCharacters(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
		// log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	standard := strings.Split(string(content), "\n")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	return standard, nil
}

func convert(banner []string, text string) string {
	lines := strings.Split(text, "\\n")
	result := ""
	for index, line := range lines {
		result += convertLine(banner, line)
		if index != len(lines) {
			result += "\n"
		}
	}
	return result
}

func convertLine(banner []string, line string) string {
	result := ""
	for index := 1; index <= 8; index++ {
		for _, char := range []rune(line) {
			// if !(char >= 32 && char <= 126) {
			// 	return "Error: Unsupported Characters: " + line
			// }
			result += banner[(int(char)-32)*9+index]
		}
		if index != 8 {
			result += "\n"
		}
	}
	return result
}

func inputValidation(input string) bool {
	char := map[int]bool{}
	for i := 0; i <= 94; i++ {
		char[(i + 32)] = true
	}

	for _, key := range input {

		if !char[int(key)] {
			fmt.Print("Not valid character: ", string(key), "\n")
			return false
		}
	}
	return true
}

func createFile(fileName string) {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func writeFile(str string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	defer file.Close()
	if isError(err) {
		return
	}

	file.Truncate(0)
	file.Seek(0, 0)

	_, err = file.WriteString(str)
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
