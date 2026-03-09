package main

import (
	"strings"
)

func getInputMatrix(str string) [][]string {
	r := strings.Split(strings.TrimSpace(str), "\n")

	rows := len(r)
	cols := len(r[0])
	matrix := make([][]string, rows)

	for i := range matrix {
		matrix[i] = make([]string, cols)

		for j, ch := range r[i] {
			matrix[i][j] = string(ch)
		}
	}

	return matrix
}

var inputStr = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + 
`
