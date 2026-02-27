package main

import (
	"fmt"
	"strconv"
)

/*
- iterate over batteries in bank twice
- first pass: iterate over all batteries in bank to find largest
- second pass: iterate over batteries from largest to end of string (bank)
- handle edge case where largest is in last position
*/

func main() {
	count := 0

	input := parseInputString(inputStr)
	for _, bank := range input {
		batt1Idx := getBatt(bank)
		batt2Idx := getBatt(bank[batt1Idx+1:])

		batt1 := int(bank[batt1Idx] - '0')
		batt2 := int(bank[batt1Idx+1+batt2Idx] - '0')
		battEnd := int(bank[len(bank)-1] - '0')

		if battEnd > batt2 {
			n, _ := strconv.Atoi(fmt.Sprintf("%d%d", batt1, battEnd))
			count += n
		} else {
			n, _ := strconv.Atoi(fmt.Sprintf("%d%d", batt1, batt2))
			count += n
		}
	}

	fmt.Println(count)
}

func getBatt(bank string) int {
	currMaxIdx := 0

	for i := range len(bank) - 1 {
		if int(bank[i]-'0') > int(bank[currMaxIdx]-'0') {
			currMaxIdx = i
		}
	}

	return currMaxIdx
}
