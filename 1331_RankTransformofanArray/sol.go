package arrayRankTransform

import "sort"

func arrayRankTransform(arr []int) []int {
	return solution(arr)
}

func solution(arr []int) []int {

	if len(arr) == 0 {
		return []int{}
	}

	nonDup := []int{}
	nonDup = append(nonDup, arr[0])

	for i := 0; i < len(arr); i++ {
		check := true
		tempNum := 0
		for j := 0; j < len(nonDup); j++ {
			tempNum = arr[i]
			if nonDup[j] == arr[i] {
				check = false
				break
			} else {
				check = true
			}
		}
		if check {
			nonDup = append(nonDup, tempNum)
		}
	}

	// sort

	sort.Ints(nonDup)

	// map index
	ans := []int{}

	for i := 0; i < len(arr); i++ {
		index := sort.SearchInts(nonDup, arr[i])
		ans = append(ans, index+1)
	}

	return ans

}
