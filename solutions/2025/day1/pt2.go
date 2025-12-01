package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

// NOT WORKING
func pt2(input string) (string, error) {
	cnt := 0
	lock := 50
	lines := strings.Split(input, "\n")

	for i := range lines {
		line := lines[i]
		c, j := utf8.DecodeRuneInString(line)
		num, err := strconv.Atoi(line[j:])
		if err != nil {
			panic(err)
		}

		fmt.Println("Lock: ", lock)
		fmt.Println("Input: ", line)
		switch c {
		case 'R':
			val := lock + num
			if val >= 100 {
				fmt.Println("    Crossed zero (overflow)")
				oldCnt := cnt
				cnt += val / 100
				fmt.Println("    Count:", cnt, "(", oldCnt, "+", val/100, ")")
			}

			lock = val % 100
		case 'L':
			val := lock - num
			if val < 0 {
				fmt.Println("    Crossed zero (underflow)")
				cnt += 1 - val/100
				fmt.Println("    Count:", cnt)
				lock = ((val % 100) + 100) % 100
			} else {
				lock = val
			}
		}

		fmt.Println("#################")
	}

	return strconv.Itoa(cnt), nil
}
