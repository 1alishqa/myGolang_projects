package student

func SplitWhiteSpaces(str string) []string {
	length := 0
	strLen := 0
	start := 0
	j := 0
	wasLetter := false

	for _, v := range str {

		if (v == '\n' || v == '\t' || v == ' ') && wasLetter {
			length++
			wasLetter = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			wasLetter = true
		}

		strLen++
	}

	if wasLetter {
		length++
		wasLetter = false
	}

	arr := make([]string, length)

	for i, v := range str {

		if (v == '\n' || v == '\t' || v == ' ') && wasLetter {
			arr[j] = str[start:i]
			j++
			wasLetter = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			if !wasLetter {
				start = i
			}
			wasLetter = true
		}

	}

	if wasLetter {
		arr[j] = str[start:strLen]
	}

	return arr
}