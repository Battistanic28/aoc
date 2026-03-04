package main

/*
Steps:
- Parse input into slice of ranges and slice of ingredients
- Define Range type with isFresh method - ingredient is fresh if within range
- Loop through indredients and check for presence in each range
- This is the naive approach - consider optimizations

Gotchas:
- Break out of loop once ingredient is identified as not fresh so as to not double count
*/

import "fmt"

type Range struct {
	Start int64
	End   int64
}

func (r Range) isFresh(id int64) bool {
	if id >= r.Start && id <= r.End {
		return true
	}

	return false
}

func main() {
	count := 0
	ranges, ingredients := parseInput(inputStr)

outer:
	for _, i := range ingredients {
		for _, r := range ranges {
			isFresh := r.isFresh(int64(i))

			if isFresh {
				count += 1
				continue outer
			}
		}
	}

	fmt.Println(count)
}
