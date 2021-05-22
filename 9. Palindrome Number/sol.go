func isPalindrome(x int) bool {
	str1 := strconv.Itoa(x)
	str2 := ""

	for i := len(str1) - 1; i >= 0; i-- {
		str2 += string(str1[i])
	}

	return str1 == str2
}