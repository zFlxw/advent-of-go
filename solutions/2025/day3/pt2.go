package day3

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	sum := 0
	banks := strings.SplitSeq(input, "\n")
	for bank := range banks {
		val, err := strconv.Atoi(highestJoltage(bank, 12))
		if err != nil {
			return "", err
		}

		sum += val
	}

	return strconv.Itoa(sum), nil
}
