package main

import (
	"fmt"
	student ".."
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)

	fmt.Println(result)

	result2 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", 
	"five", "fingers", "of", "each", "hand"}
	student.AdvancedSortWordArr(result2, student.Compare)

	fmt.Println(result2)

	result3 := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	student.AdvancedSortWordArr(result3, student.Compare)

	fmt.Println(result3)

	result4 := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	student.AdvancedSortWordArr(result4, Compare)

	fmt.Println(result4)

	result5 := []string{"The", "computing", "consisted", "device", "each", 
	"earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	student.AdvancedSortWordArr(result5, Compare)

	fmt.Println(result5)

	result6 := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	student.AdvancedSortWordArr(result6, Compare)

	fmt.Println(result6)
}

func Compare(b, a string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}