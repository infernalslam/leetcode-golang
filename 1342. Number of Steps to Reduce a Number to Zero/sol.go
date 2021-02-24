func numberOfSteps(num int) int {
	return sol(num)
}

func sol(num int) int {

	ans := 0
	for num > 0 {
		if isOdd(num) {
			num -= 1
		} else {
			num /= 2
		}

		ans++

	}

	return ans
}

func isOdd(num int) bool {
	return num%2 != 0
}

