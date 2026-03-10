package main

import (
	"fmt"
	"strconv"
)

type Calc struct {
	result int
}

func main() {
	m, r, c := getInputMatrix(inputStr)
	operators := m[r-1]
	result := 0

	for col := range c {
		s := getColumnSlice(m, col)
		operator := operators[col]
		result += calculate(s, operator)
	}

	fmt.Println(result)
}

func getColumnSlice(matrix [][]string, col int) []int {
	s := []int{}
	for row := range len(matrix) - 1 {
		num, _ := strconv.Atoi(matrix[row][col])
		s = append(s, num)
	}

	return s
}

func calculate(nums []int, operator string) int {
	c := &Calc{}
	switch operator {
	case "*":
		c.mult(nums)
	case "+":
		c.add(nums)
	}

	return c.result
}

func (c *Calc) mult(nums []int) {
	c.result = 1
	for _, n := range nums {
		c.result = c.result * n
	}
}

func (c *Calc) add(nums []int) {
	c.result = 0
	for _, n := range nums {
		c.result = c.result + n
	}
}
