func judgeCircle(moves string) bool {
	return sol(moves)
}

func sol(moves string) bool {

	str := []rune(moves)

	x := 0
	y := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "U" {
			y++
		} else if string(str[i]) == "D" {
			y--
		} else if string(str[i]) == "L" {
			x--
		} else if string(str[i]) == "R" {
			x++
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}

} 