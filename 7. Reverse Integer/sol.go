func reverse(x int) int {
	nums := strconv.Itoa(x)

	str := ""

	for i := len(nums) - 1; i >= 0; i-- {
		if string(nums[i]) == "-" {
			str = "-" + str
			break
		}
		str += string(nums[i])
	}

	i, _ := strconv.Atoi(str)

	// wtf code
	if i > 2147483647 {
		return 0
	}
	if i < -2147483648 {
		return 0
	}
	//
	return i
}
