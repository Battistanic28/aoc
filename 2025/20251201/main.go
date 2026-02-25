package main

// RUN: go run main.go input.go

/*
Steps:
- Parse input (ex. L10, R15 --> [-10, 15])
- Iterate over input values, adjust dial for each i in range(input)
- Set dial to new current position
- if currPosition == 0, increment count

Gotchas:
- Handle dial adjustments with value >= incrementCount (+/-)
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	d := Dial{
		currPosition:   50,
		incrementCount: 99,
	}

	currPosition := d.currPosition
	count := 0

	parsedInput := parseInputString(sampleStr)

	for i := range parsedInput {
		currPosition = d.Adjust(parsedInput[i])
		d.currPosition = currPosition

		if currPosition == 0 {
			count++
		}
	}

	fmt.Println(count)
}

type Dial struct {
	currPosition   int
	incrementCount int
}

func (d Dial) Adjust(increment int) int {
	position := (d.currPosition + increment) % d.incrementCount

	return position
}

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
