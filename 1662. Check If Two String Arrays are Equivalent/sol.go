func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return sol(word1, word2)
}

func sol(w1, w2 []string) bool {
	s1 := ""
	s2 := ""

	for _, val := range w1 {
		s1 += string(val)
	}

	for _, val := range w2 {
		s2 += string(val)
	}

	if s1 == s2 {
		return true
	}

	return false

}