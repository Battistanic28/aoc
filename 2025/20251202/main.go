package main

import "fmt"

/*
- parse input into distinct Range instances
- define validation algorithm
- iterate over value in Range
- run validation algorithm for each value in Range
*/

func main() {
	count := 0
	for _, r := range parsedInput {
		fmt.Println("range", r)
		for j := r.Start; j <= r.End; j++ {
			validate(j)
		}
	}

	fmt.Println(count)
}

func validate(number int64) int64 {
	fmt.Println("validating", number)
	return number // placeholder - if invalid return number, else return 0
}
