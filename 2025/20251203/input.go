package main

import (
	"strings"
)

func parseInputString(str string) []string {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(str), "\n")
	var s []string

	for _, line := range lines {
		if line == "" {
			continue
		}
		s = append(s, line)
	}

	return s
}

var inputStr = `
987654321111111
811111111111119
234234234234278
818181911112111
`
