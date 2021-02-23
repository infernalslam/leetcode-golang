func runningSum(nums []int) []int {
	return sum(nums)
}

func sum(nums []int) []int {
	temp := 0
	var sum []int
	for i, n := range nums {
		if i == 0 {
			temp = n
		} else {
			temp += n
		}
		sum = append(sum, temp)
	}
	return sum
}

