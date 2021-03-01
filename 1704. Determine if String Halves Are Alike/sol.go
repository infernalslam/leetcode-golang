func halvesAreAlike(s string) bool {
	return sol(s)
}

func sol(s string) bool {

	dl := len(s) / 2

	s1 := 0
	s2 := 0

	for i := 0; i < dl; i++ {
		s1 += checkVowel(s[i])
	}

	for i := dl; i < len(s); i++ {
		s2 += checkVowel(s[i])
	}

	return s1 == s2
}

func checkVowel(s byte) int {

	vowelU := []string{"a", "e", "i", "o", "u"}
	vowelL := []string{"A", "E", "I", "O", "U"}

	sum := 0

	for i := 0; i < 5; i++ {
		if string(s) == vowelU[i] {
			sum += 1
			continue
		} else if string(s) == vowelL[i] {
			sum += 1
			continue
		}
	}

	return sum

}