func restoreString(s string, indices []int) string {
	return sol(s, indices)
}

func sol(s string, indices []int) string {

	rune := make([]rune, len(indices))

	for index, w := range s {
		rune[indices[index]] = w
	}

	fmt.Println(string(rune))

	return string(rune)
}