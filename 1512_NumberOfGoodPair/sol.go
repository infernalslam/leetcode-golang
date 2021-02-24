package numberOfGoodPair

func numIdenticalPairs(nums []int) int {
	return test(nums)
}

func test(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				sum += 1
			}
		}
	}
	return sum
}
