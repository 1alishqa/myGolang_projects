package student

func AdvancedSortWordArr(array []string, f func(a, b string) int) {

	for i := 0; i < len(array); i++ {

		for j := len(array) - 1; j > i; j-- {

			if f(array[j], array[j-1]) < 0 {

				Swap(&array[j], &array[j-1])

			}

		}
	}
}

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func Swap(a, b *string) {
	x := *a
	y := *b
	*a = y
	*b = x
}