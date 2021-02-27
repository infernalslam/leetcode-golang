func findNumbers(nums []int) int {
	return sol(nums)
}

func sol(nums []int) int {

	ans := 0
	for _, n := range nums {

		if isEven(n) {
			if isCheckDigitEven(n) {
				ans++
			}
			continue
		}

		if isCheckDigitEven(n) {
			ans++
		}

	}

	return ans
}

func isEven(n int) bool {
	return n%2 == 0
}

func isCheckDigitEven(n int) bool {
	s := strconv.Itoa(n)
	le := len(s)
	check := false
	if le%2 == 0 {
		check = true
	}
	return check
}