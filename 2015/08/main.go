package main

import (
	"fmt"
	"strings"

	"github.com/rdoois/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput()
	fmt.Println("Answer 1:", first(input))
	fmt.Println("Answer 2:", second(input))
}

func first(input string) int {
	var lines = strings.Split(input, "\n")
	var literals, chars int

	for _, line := range lines {
		literals += len(line)

		for i := 1; i < len(line)-1; i++ {
			curr := line[i]
			if curr == '\\' {
				next := line[i+1]
				if next == '\\' || next == '"' {
					i++
				} else if next == 'x' {
					i += 3
				}
			}
			chars++
		}
	}

	return literals - chars
}

func second(input string) int {
	var lines = strings.Split(input, "\n")
	var literals, encoded int

	for _, line := range lines {
		literals += len(line)
		encoded += 2

		for i := 0; i < len(line); i++ {
			curr := line[i]
			if curr == '\\' || curr == '"' {
				encoded++
			}

			encoded++
		}
	}

	return encoded - literals
}
