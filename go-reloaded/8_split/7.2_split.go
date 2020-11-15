package student

func Split(str, charset string) []string {

	if EmptyCharset(charset) {		// if charset = ""
		return SimpleArray(str)
	}

	array := make([]string, SizeofArr(str, charset))
	
	indx := 0
	x := 0
	l := Strlen(str)
	l2 := Strlen(charset)

	for i := 0; i <= l - l2; i++ {
		if str[i : i + l2] == charset {
			array[x] = str[indx : i]
			x++
			indx = i + l2
		}
	}
	array[x] = str[indx : l]

	return array
}

func SizeofArr(str, charset string) int {		// find how many elems in arr should be

	count := 1
	l := Strlen(str)
	l2 := Strlen(charset)

	for i := 0; i <= l - l2; i++ {
		if str[i : i + l2] == charset {
			count++
		}
	}
	return count
}

func EmptyCharset(charset string) bool {
	if Strlen(charset) == 0 {
		return true
	}
	return false
}

func SimpleArray(str string) []string {
	arr := make([]string, Strlen(str))

	for i, v := range str {
		arr[i] = string(v)
	}

	return arr
}

func Strlen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
