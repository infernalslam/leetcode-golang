func maximumWealth(accounts [][]int) int {
	return wealth(accounts)
}

func wealth(accounts [][]int) int {
	var accSum []int
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for _, v := range accounts[i] {
			sum += v
		}
		accSum = append(accSum, sum)
	}

	return max(accSum)
}

func max(sums []int) int {

	max := sums[0]

	for i := range sums {
		if sums[i] > max {
			max = sums[i]
		}
	}

	return max
}