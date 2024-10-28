package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rdoois/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput()
	fmt.Println("Answer 1:", first(input))
}

func first(input string) int {
	for i := 0; i >= 0; i++ {
		hexadecimal := fmt.Sprintf("%s%d", input, i)
		hashed := md5.Sum([]byte(hexadecimal))
		word := hex.EncodeToString(hashed[:])
		if strings.HasPrefix(word, "00000") {
			return i
		}
	}

	return 0
}
