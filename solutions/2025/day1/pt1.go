package day1

import (
	"advent-of-go/utils"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
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

		switch c {
		case 'R':
			lock = int(math.Abs(float64((lock + num) % 100)))
		case 'L':
			val := (lock - num) % 100
			if val < 0 {
				lock = 100 + val
			} else {
				lock = val
			}
		}

		if lock == 0 {
			cnt += 1
		}
	}

	return strconv.Itoa(cnt), nil
}
