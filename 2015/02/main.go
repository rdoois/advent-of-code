package main

import (
	"fmt"
	"strings"

	"github.com/rdoois/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Answer 1:", first(input))
}

func first(input string) int {
	var paper int
	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var lenght, width, heigth int
		fmt.Sscanf(line, "%dx%dx%d", &lenght, &width, &heigth)

		a0 := lenght * width
		a1 := lenght * heigth
		a2 := width * heigth

		paper += (a0 * 2) + (a1 * 2) + (a2 * 2) + min(a0, a1, a2)
	}

	return paper
}
