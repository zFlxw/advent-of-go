package day2

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, v := range ranges {
		parts := strings.Split(v, "-")

		first, firstErr := strconv.Atoi(parts[0])
		last, lastErr := strconv.Atoi(parts[1])
		if firstErr != nil {
			return "", firstErr
		}

		if lastErr != nil {
			return "", lastErr
		}

		for i := first; i <= last; i++ {
			maxSubstringLen := len(strconv.Itoa(i)) / 2
			if i < 10 {
				continue
			}

			leftString := strconv.Itoa(i)[:maxSubstringLen]
			rightString := strconv.Itoa(i)[maxSubstringLen:]

			if leftString[0] != '0' && leftString == rightString {
				sum += i
			}
		}
	}

	return strconv.Itoa(sum), nil
}
