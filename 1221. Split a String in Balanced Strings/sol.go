package balancedStringSplit

func balancedStringSplit(s string) int {
	return sol(s)
}

func sol(s string) int {
	str := []rune(s)
	l := 0
	r := 0
	ans := 0
	for _, v := range str {

		if "L" == string(v) {
			l++
		} else {
			r++
		}

		if l == r {
			ans++
			l = 0
			r = 0
		}

	}

	return ans
}
