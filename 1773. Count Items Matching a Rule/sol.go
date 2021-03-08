func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	return sol(items, ruleKey, ruleValue)
}

func sol(items [][]string, key, value string) int {
	sum := 0
	for i := 0; i < len(items); i++ {
		index := 0
		if key == "color" {
			index = 1
		} else if key == "type" {
			index = 0
		} else {
			index = 2
		}
		sum += check(index, value, items[i])
	}
	return sum
}

func check(index int, value string, arr []string) int {
	if arr[index] == value {
		return 1
	} else {
		return 0
	}
}