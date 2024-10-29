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
	var grid [1000][1000]bool

	for _, line := range lines {
		coords, command := getCoords(line)
		handleLights(&grid, coords, command)
	}

	return countLights(grid)
}

func second(input string) int {
	var lines = strings.Split(input, "\n")
	var grid [1000][1000]int

	for _, line := range lines {
		coords, command := getCoords(line)
		handleBrightness(&grid, coords, command)
	}

	return countBrightness(grid)
}

func getCoords(line string) ([4]int, string) {
	var coords [4]int
	var command string

	if strings.HasPrefix(line, "turn on") {
		command = "on"
	} else if strings.HasPrefix(line, "turn off") {
		command = "off"
	} else {
		command = "toggle"
	}

	words := strings.Split(line, " ")
	size := len(words)
	var start, end [2]int

	fmt.Sscanf(words[size-3], "%d,%d", &start[0], &start[1])
	fmt.Sscanf(words[size-1], "%d,%d", &end[0], &end[1])

	coords = [4]int{
		start[0], start[1], end[0], end[1],
	}

	return coords, command
}

func handleLights(grid *[1000][1000]bool, coords [4]int, command string) {
	for i := coords[0]; i <= coords[2]; i++ {
		for j := coords[1]; j <= coords[3]; j++ {
			if command == "on" {
				grid[i][j] = true
			} else if command == "off" {
				grid[i][j] = false
			} else if command == "toggle" {
				grid[i][j] = !grid[i][j]
			}
		}
	}
}

func countLights(grid [1000][1000]bool) int {
	var lights int

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				lights++
			}
		}
	}

	return lights
}

func handleBrightness(grid *[1000][1000]int, coords [4]int, command string) {
	for i := coords[0]; i <= coords[2]; i++ {
		for j := coords[1]; j <= coords[3]; j++ {
			if command == "on" {
				grid[i][j] += 1
			} else if command == "off" && grid[i][j] > 0 {
				grid[i][j] -= 1
			} else if command == "toggle" {
				grid[i][j] += 2
			}
		}
	}
}

func countBrightness(grid [1000][1000]int) int {
	var brightness int

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			brightness += grid[i][j]
		}
	}

	return brightness
}
