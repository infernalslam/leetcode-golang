func checkIfPangram(sentence string) bool {
	// 2847 unicode a-z
	chars := []int{}

	for _, val := range sentence {
		i, _ := strconv.Atoi(fmt.Sprintf("%d", val))

		isPush := true
		for _, jj := range chars {
			if i == jj {
				isPush = false
				break
			}
		}

		if isPush && i != 0 {
			chars = append(chars, i)
		}

	}

	sum := 0
	for _, val := range chars {
		sum += val
	}

	return sum == 2847
}