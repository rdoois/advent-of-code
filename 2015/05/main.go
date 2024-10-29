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
	var words int
	var lines = strings.Split(input, "\n")

	for _, word := range lines {
		if containsAtLeastThreeVowels(word) &&
			containsAtLeastTwoEqualLettersInARow(word) &&
			doesNotContainTheStrings(word) {
			words++
		}
	}

	return words
}

func second(input string) int {
	var words int
	var lines = strings.Split(input, "\n")

	for _, word := range lines {
		if containsAPairAtLeastTwice(word) && containsAtLeastALetterBetweenAPair(word) {
			words++
		}
	}

	return words
}

func containsAtLeastThreeVowels(word string) bool {
	var vowels int

	for _, r := range word {
		if strings.ContainsRune("aeiou", r) {
			vowels++
			if vowels == 3 {
				return true
			}
		}
	}

	return false
}

func containsAtLeastTwoEqualLettersInARow(word string) bool {
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			return true
		}
	}

	return false
}

func doesNotContainTheStrings(word string) bool {
	var forbiddens = [4]string{"ab", "cd", "pq", "xy"}

	for _, f := range forbiddens {
		if strings.Contains(word, f) {
			return false
		}
	}

	return true
}

func containsAPairAtLeastTwice(word string) bool {
	var pair string
	var pairs = make(map[string]bool)

	for i := 1; i < len(word); i++ {
		if i > 1 && word[i] == word[i-1] && word[i] == word[i-2] {
			continue
		}
		pair = fmt.Sprintf("%c%c", word[i-1], word[i])
		if pairs[pair] {
			return true
		}

		pairs[pair] = true
	}

	return false
}

func containsAtLeastALetterBetweenAPair(word string) bool {
	for i := 2; i < len(word); i++ {
		if word[i-2] == word[i] {
			return true
		}
	}

	return false
}
