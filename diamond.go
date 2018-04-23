package diamond

import "fmt"

/*
Diamond ...
*/
func Diamond(letter rune) string {
	if letter == 'A' {
		return "A"
	}

	ret := "A\n"

	for _, l := range "BC" {
		if letter == l {
			break
		}
		ret = ret + fmt.Sprintf("%c %c\n", l, l)
	}
	ret = ret + fmt.Sprintf("%c  %c\n", letter, letter)
	for _, l := range "BC" {
		if letter == l {
			break
		}
		ret = ret + fmt.Sprintf("%c%c\n", l, l)
	}
	return ret + "A"
}
