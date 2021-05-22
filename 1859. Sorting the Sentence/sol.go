func sortSentence(s string) string {
	sentence := strings.Split(s, " ")
	index := []Index{}
	for _, val := range sentence {
		chars := string(val)
		temp := ""
		for i := 0; i < len(chars); i++ {
			if chars[i] >= 49 && chars[i] <= 57 {
				s := string(chars[i])
				i, _ := strconv.Atoi(s)
				index = append(index, Index{Index: i - 1, Text: temp})
			} else {
				temp += string(chars[i])
			}
		}
	}

	ans := make([]string, len(index))

	for _, val := range index {
		ans[val.Index] = val.Text
	}

	resp := ""
	for _, val := range ans {
		resp += val + " "
	}

	return strings.TrimSpace(resp)
}

type Index struct {
	Index int
	Text  string
}
