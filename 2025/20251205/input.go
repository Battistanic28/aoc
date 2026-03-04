package main

import (
	"strconv"
	"strings"
)

func parseInput(inputStr string) ([]Range, []int64) {
	data := strings.Split(strings.TrimSpace(inputStr), "\n\n")
	idRanges := makeRanges(data[0])
	ingredients := getIngredients(data[1])

	return idRanges, ingredients
}

func makeRanges(rangeStr string) []Range {
	lines := strings.Split(strings.TrimSpace(rangeStr), "\n")
	var s []Range

	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)

		r := Range{
			Start: start,
			End:   end,
		}

		s = append(s, r)
	}

	return s
}

func getIngredients(str string) []int64 {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	var s []int64

	for _, line := range lines {
		num, _ := strconv.ParseInt(line, 10, 64)

		s = append(s, num)
	}

	return s
}

var inputStr = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`
