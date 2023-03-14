package dfutil

import (
	"errors"
	"strings"
)

var dictionary = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var order = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

// Itor converts an interger to a roman numeral
func Itor(n int) (string, error) {
	if n > 3999 {
		return "", errors.New("the integer provided is too big")
	}

	var s strings.Builder

	for _, i := range order {
		for n >= i {
			n -= i
			s.WriteString(dictionary[i])
		}
	}
	return s.String(), nil
}
