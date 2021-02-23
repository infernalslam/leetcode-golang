func search(nums []int, target int) int {
	return binaryTree(nums, target)
}

func binaryTree(nums []int, target int) int {

	index := -1
	for i, n := range nums {
		if target == n {
			index = i
			break
		}
	}

	return index
}