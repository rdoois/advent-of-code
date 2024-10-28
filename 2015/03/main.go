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

func second(input string) int {
	houses := make(map[[2]int]bool)
	var move, santa, robot [2]int
	houses[santa] = true

	for i, r := range input {
		move = directions[r]
		if i%2 == 0 {
			santa = [2]int{santa[0] + move[0], santa[1] + move[1]}
			houses[santa] = true
		} else {
			robot = [2]int{robot[0] + move[0], robot[1] + move[1]}
			houses[robot] = true
		}
	}

	return len(houses)
}
