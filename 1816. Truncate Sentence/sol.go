func truncateSentence(s string, k int) string {
	text := strings.Split(s, " ")
	fmt.Println(text)
	ans := ""
	for i, val := range text {
		if i == k {
			break
		}
		ans += val + " "
	}
	return strings.TrimSpace(ans)
}