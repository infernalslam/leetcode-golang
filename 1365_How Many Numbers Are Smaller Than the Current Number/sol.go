package smallernumbersthancurrent

func smallerNumbersThanCurrent(nums []int) []int {
	return sol(nums)
}

func sol(nums []int) []int {

	ans := []int{}
	for _, n := range nums {

		loop := 0
		temp := 0

		for loop < len(nums) {
			if n > nums[loop] {
				temp++
			}
			loop++
		}

		ans = append(ans, temp)
	}

	return ans
}
