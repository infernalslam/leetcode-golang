func sortArrayByParity(A []int) []int {
	return sol(A)
}

func sol(a []int) []int {

	position := 0

	for i := 0; i < len(a); i++ {
		// first loop : [3, 1, 2, 4]
		if a[i]%2 == 0 {
			// if even number 2, 4 ..
			// swap array
			// position = 0, i = 2
			// first (2) => a[0] = a[2] , a[2] = a[0] => [2, 1, 3, 4]
			// second (4) => a[1] = a[3], a[3] = a[1] => [2, 4, 3, 1]
			a[position], a[i] = a[i], a[position]
			position++
		}
	}

	return a
}
