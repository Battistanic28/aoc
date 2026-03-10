package main

import (
	"strings"
)

func getInputMatrix(str string) ([][]string, int, int) {
	r := strings.Split(strings.TrimSpace(str), "\n")

	rows := len(r)
	cols := len(strings.Fields(r[0]))
	matrix := make([][]string, rows)

	for i := range matrix {
		matrix[i] = make([]string, cols)
		fields := strings.Fields(r[i])

		for j, ch := range fields {
			matrix[i][j] = string(ch)
		}
	}

	return matrix, rows, cols
}

var inputStr = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + 
`
