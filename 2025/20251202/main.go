package main

import (
	"fmt"
	"strconv"
)

/*
- parse input into distinct Range instances
- define validation algorithm
- iterate over value in Range
- run validation algorithm for each value in Range
*/

func main() {
	var count int64 = 0
	for _, r := range parsedInput {
		for j := r.Start; j <= r.End; j++ {
			value := validate(j)
			count += value
		}
	}

	fmt.Println(count)
}

func validate(num int64) int64 {
	str := strconv.FormatInt(num, 10)
	len := len(str)
	mid := len / 2
	firstHalf := str[:mid]
	secondHalf := str[mid:]
	if firstHalf == secondHalf {
		return num
	}

	return 0
}
