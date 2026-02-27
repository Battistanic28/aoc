package main

import (
	"strconv"
	"strings"
)

func parseInputString(str string) []int {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(str), "\n")
	var s []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		dir := string(line[0])
		dist, _ := strconv.Atoi(line[1:])

		if dir == "L" {
			s = append(s, 0-dist)

		}

		if dir == "R" {
			s = append(s, dist)

		}
	}

	return s
}

var inputStr string = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`
