package day3

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	sum := 0
	banks := strings.SplitSeq(input, "\n")
	for bank := range banks {
		val, err := strconv.Atoi(highestJoltage(bank, 2))
		if err != nil {
			return "", err
		}

		sum += val
	}

	return strconv.Itoa(sum), nil
}

func highestJoltage(s string, n int) string {
	stack := make([]byte, 0, len(s))
	remove := len(s) - n

	for i := 0; i < len(s); i++ {
		c := s[i]
		for len(stack) > 0 && c > stack[len(stack)-1] && remove > 0 {
			stack = stack[:len(stack)-1]
			remove--
		}

		stack = append(stack, c)
	}

	return string(stack[:n])
}
