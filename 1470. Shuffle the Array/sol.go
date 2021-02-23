func shuffle(nums []int, n int) []int {
	return quiz(nums, n)
}

func quiz(nums []int, n int) []int {

	a1 := nums[0:n]
	a2 := nums[n:len(nums)]

	s := []int{}

	for i := 0; i < n; i++ {
		s = append(s, a1[i])
		s = append(s, a2[i])
	}

	return s
}
