package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	minus := Minus(nbr)
	if minus {
		v := nbr[1:]
		if !Valid(v, baseFrom) {
			return "NV"
		}
	} else if !Valid(nbr, baseFrom) {
		return "NV"
	}
    n := AtoiBase(nbr, baseFrom)
    return PrintNbrBase(n, baseTo)
}

func AtoiBase(s string, base string) int {
    lenBase := count(base)
    arrBase := []rune(base)
    arrStr := []rune(s)
    ans := 0
    for _, i := range arrStr {
        for index, j := range arrBase {
            if i == j {
                ans = lenBase*ans + index
            }
        }
    }
    return ans
}

func PrintNbrBase(nbr int, base string) string {
    baseLen := count(base)
    baseArr := []rune(base)
    ans := ""
    for nbr != 0 {
        tmp := 0
        if nbr < 0 {
            tmp = -nbr % baseLen
        } else {
            tmp = nbr % baseLen
        }
        if tmp > 0 {
            ans += string(baseArr[tmp])
        } else if tmp == 0 {
            ans += string(baseArr[0])
        } else {
            ans += string(baseArr[-tmp])
        }
        nbr = nbr / baseLen
    }
    ans = Reverse(ans)
    return ans
}

func count(s string) int {
    cnt := 0
    for range s {
        cnt++
    }
    return cnt
}


func Reverse(s string) string {
    ans := ""
    for _, v := range s {
        ans = string(v) + ans
    }
    return ans
}


func Valid(n, base string) bool {
	var eq int
	
	for i := range n {

		eq = 0
		for l := range base {
			if n[i] == base[l] {
				eq++
			}
		}

		if eq == 0 {
			return false
		}
	}
	return true
}

func Minus(s string) bool {
	if s[0] == '-' {
		return true
	}
	return false
}
