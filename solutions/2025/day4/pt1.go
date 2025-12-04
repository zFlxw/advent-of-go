package day4

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	neighbours := make([][]int, len(lines))
	for i := range neighbours {
		neighbours[i] = make([]int, len(lines[0]))
		for j := range neighbours[i] {
			neighbours[i][j] = -1
		}
	}

	for i, line := range lines {
		for j, char := range line {
			if char != '@' {
				continue
			}

			count := 0

			for _, dir := range directions {
				row, col := i+dir[0], j+dir[1]
				if row >= 0 && row < len(lines) && col >= 0 && col < len(line) {
					if lines[row][col] == '@' {
						count++
					}
				}
			}

			neighbours[i][j] = count
		}
	}

	amount := 0
	for row := range neighbours {
		for col := 0; col < len(neighbours[row]); col++ {
			val := neighbours[row][col]
			if val >= 0 && val < 4 {
				amount++
			}
		}
	}

	return strconv.Itoa(amount), nil
}
