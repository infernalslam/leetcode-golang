func createTargetArray(nums []int, index []int) []int {
  
  res := make([]int, 0, len(nums))

	for i, v := range index {
		res = append(res[:v], append([]int{nums[i]}, res[v:]...)...)
	}
  
  
  return res
  
}