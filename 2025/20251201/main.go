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
)

func main() {
	d := Dial{
		currPosition:   50,
		incrementCount: 100, // 99, shift by 1
	}

	currPosition := d.currPosition
	count := 0

	parsedInput := parseInputString(inputStr)

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

	if position < 0 {
		position += d.incrementCount
	}

	return position
}
