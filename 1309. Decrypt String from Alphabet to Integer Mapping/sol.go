func freqAlphabets(s string) string {
	return sol(s)
}

func sol(s string) string {
	temp := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "#" {
			// 10 + 96  =  106 => j
			// 11 + 96  = 107 => k
			// 12 + 96 = 108 => l
			// string(s[i - 2]) + string(s[i - 1]) // 26 + 96 => 122 => z
			num, _ := strconv.Atoi(string(s[i-2]) + string(s[i-1]))
			temp = append(temp, string(num+96))
			i -= 2
		} else {
			num, _ := strconv.Atoi(string(s[i]))
			temp = append(temp, string(num+96))
		}
	}

	ans := pasreAnswer(temp)
	return ans
}

func pasreAnswer(s []string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += s[i]
	}
	return str
}