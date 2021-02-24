package numJewelsInStones

func numJewelsInStones(jewels string, stones string) int {
	return sol(jewels, stones)
}

func sol(jewels string, stones string) int {

	ans := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				ans++
			}
		}
	}

	return ans
}
