package main

import (
	"fmt"

	"github.com/rdoois/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Answer 1:", first(input))
}

var directions = map[rune][2]int{
	'^': {0, 1},
	'v': {0, -1},
	'>': {1, 0},
	'<': {-1, 0},
}

func first(input string) int {
	houses := make(map[[2]int]bool)
	var move, santa [2]int
	houses[santa] = true
	for _, r := range input {
		move = directions[r]
		santa = [2]int{santa[0] + move[0], santa[1] + move[1]}
		houses[santa] = true
	}

	return len(houses)
}
