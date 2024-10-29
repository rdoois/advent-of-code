package main

import (
	"fmt"
	"math"
	"strconv"
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

	var graph = make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		graph[parts[1]] = parts[0]
	}

	return dfs(graph, "a", map[string]int{})
}

func second(input string) int {
	var lines = strings.Split(input, "\n")

	var graph = make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		graph[parts[1]] = parts[0]
	}

	graph["b"] = strconv.Itoa((dfs(graph, "a", map[string]int{})))
	return dfs(graph, "a", map[string]int{})
}

func dfs(graph map[string]string, key string, circuit map[string]int) int {
	if val, ok := circuit[key]; ok {
		return val
	}

	num, err := strconv.Atoi(key)
	if err == nil {
		return num
	}

	cmd := strings.Split(graph[key], " ")
	if len(cmd) == 1 {
		circuit[key] = dfs(graph, cmd[0], circuit)
	} else if cmd[0] == "NOT" {
		circuit[key] = (math.MaxUint16) ^ dfs(graph, cmd[1], circuit)
	} else if cmd[1] == "AND" {
		circuit[key] = dfs(graph, cmd[0], circuit) & dfs(graph, cmd[2], circuit)
	} else if cmd[1] == "OR" {
		circuit[key] = dfs(graph, cmd[0], circuit) | dfs(graph, cmd[2], circuit)
	} else if cmd[1] == "LSHIFT" {
		circuit[key] = dfs(graph, cmd[0], circuit) << dfs(graph, cmd[2], circuit)
	} else if cmd[1] == "RSHIFT" {
		circuit[key] = dfs(graph, cmd[0], circuit) >> dfs(graph, cmd[2], circuit)
	}

	return circuit[key]
}
