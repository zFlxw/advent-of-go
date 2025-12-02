package day2

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
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
			numStr := strconv.Itoa(i)
			numLen := len(numStr)
			if i < 10 {
				continue
			}

			for j := 1; j <= numLen/2; j++ {
				if numLen%j == 0 {
					pattern := numStr[:j]
					repeatedPattern := strings.Repeat(pattern, numLen/j)
					if repeatedPattern == numStr && pattern[0] != '0' {
						sum += i
						break
					}
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}
