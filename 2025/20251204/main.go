package main

import "fmt"

func main() {
	count := 0
	m := getInputMatrix(inputStr)
	maxAdjacentRollCount := 4

	rows := len(m)
	cols := len(m[0])

	for i := range rows {
		for j := range cols {
			if m[i][j] == "@" {
				adjCount := adjacentRollCount(m, i, j)

				if adjCount < maxAdjacentRollCount {
					count += 1
				}
			}
		}
	}

	fmt.Println(count)
}

func adjacentRollCount(matrix [][]string, rowPosition int, colPosition int) int {
	count := 0
	rows := len(matrix)
	cols := len(matrix[0])
	offsets := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, o := range offsets {
		r, c := rowPosition+o[0], colPosition+o[1]

		if r >= 0 && r < rows && c >= 0 && c < cols {
			if matrix[r][c] == "@" {
				count += 1
			}
		}

	}

	return count
}
