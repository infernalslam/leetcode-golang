func sumOfUnique(nums []int) int {
	return sol(nums)
}

func sol(nums []int) int {

	sum := 0
	for i := 0; i < len(nums); i++ {
		dupCount := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				dupCount++
			}
		}
		if dupCount == 1 {
			sum += nums[i]
		}
	}
	return sum
}