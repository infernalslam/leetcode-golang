func mergeAlternately(word1 string, word2 string) string {
	return sol(word1, word2)
}

func sol(w1, w2 string) string {
	str := ""

	maxLen := 0
	if len(w1) >= len(w2) {
		maxLen = len(w1)
	} else {
		maxLen = len(w2)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(w1) {
			str += fmt.Sprintf("%c", w1[i])
		}
		if i < len(w2) {
			str += fmt.Sprintf("%c", w2[i])
		}
	}

	return str
}