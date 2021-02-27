package interpret

import "fmt"

func interpret(command string) string {
	return sol(command)
}

func sol(command string) string {

	str := ""
	check := true

	for i := 0; i < len(command); i++ {

		rune1 := fmt.Sprintf("%c", command[i])
		if "(" == rune1 {
			rune2 := fmt.Sprintf("%c", command[i+1])
			if ")" == rune2 {
				str += "o"
				check = false
			} else {
				check = true
			}
		} else {
			if check == false {
				check = true
				continue
			} else if check == true && rune1 != ")" {
				str += rune1
			}

		}

	}

	return str
}
