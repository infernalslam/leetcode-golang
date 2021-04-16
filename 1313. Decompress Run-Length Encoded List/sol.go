func decompressRLElist(nums []int) []int {
	return sol(nums)
}

func sol(nums []int) []int {

	mapNums := make([][]int, len(nums)/2)

	for i := 0; i < len(nums); i += 2 {
		x := []int{}
		x = append(x, nums[i])
		x = append(x, nums[i+1])
		mapNums = append(mapNums, x)
	}

	ans := []int{}

	for i := 0; i < len(mapNums); i++ {
		if len(mapNums[i]) > 0 {
			for j := 0; j < mapNums[i][0]; j++ {
				ans = append(ans, mapNums[i][1])
			}
		}
	}

	fmt.Println(ans)

	return ans
}