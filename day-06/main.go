package main

import (
	"fmt"

	"github.com/noodlensk/advent-of-code-2022/util"
)

func Solution(input string) int {
	for i := 0; i < len(input)-4; i++ {
		if isAllCharsUniq(input[i : i+4]) {
			return i + 4
		}
	}

	return 0
}

func Solution2(input string) int {
	for i := 0; i < len(input)-14; i++ {
		if isAllCharsUniq(input[i : i+14]) {
			return i + 14
		}
	}

	return 0
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines[0]))
	fmt.Println("Result #2: ", Solution2(lines[0]))
}

func isAllCharsUniq(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}
