package main

import (
	"fmt"

	"github.com/rdoois/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Answer 1:", first(input))
	fmt.Println("Answer 2:", second(input))
}

func first(input string) int {
	var floor int

	for _, r := range input {
		if r == '(' {
			floor += 1
		} else if r == ')' {
			floor -= 1
		}
	}

	return floor
}

func second(input string) int {
	var floor int

	for i, r := range input {
		if r == '(' {
			floor += 1
		} else if r == ')' {
			floor -= 1
		}

		if floor == -1 {
			return i + 1
		}
	}

	return 0
}
