package thousandSeparator

import (
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	return solution(n)
}

func solution(n int) string {

	rune := strconv.Itoa(n)

	str := ""
	temp := 0

	for i := len(rune) - 1; i >= 0; i-- {
		str += fmt.Sprintf("%c", rune[i])
		temp++
		if temp == 3 {
			if i == 0 {
				break
			}
			str += "."
			temp = 0
		}
	}

	ans := ""

	for i := len(str) - 1; i >= 0; i-- {
		ans += fmt.Sprintf("%c", str[i])
	}

	return ans

}
