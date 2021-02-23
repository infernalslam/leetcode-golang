func defangIPaddr(address string) string {
	return covertString(address)
}

func covertString(str string) string {
	temp := ""
	for _, s := range str {
		rune := fmt.Sprintf("%c", s)

		if rune == "." {
			temp += `[.]`
		} else {
			temp += rune
		}
	}

	return temp
}