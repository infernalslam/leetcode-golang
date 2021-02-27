package toLowerCase

func toLowerCase(str string) string {
	return sol(str)
}

func sol(str string) string {
	text := []byte(str)
	ans := ""

	for i := 0; i < len(text); i++ {
		if text[i] >= 65 && text[i] <= 90 {
			b := text[i] + 32
			ans += string(b)
		} else {
			ans += string(text[i])
		}
	}

	return ans
}
