package main

import (
	"fmt"
	"sort"

	"github.com/noodlensk/advent-of-code-2022/util"
)

func Solution(input []string) int {
	topCalories := caloriesSorted(input)

	return topCalories[len(topCalories)-1]
}

func Solution2(input []string) int {
	top := caloriesSorted(input)

	return top[len(top)-1] + top[len(top)-2] + top[len(top)-3]
}

func caloriesSorted(input []string) []int {
	var (
		totalCalories      []int
		currentElfCalories int
	)

	for _, l := range input {
		if l == "" { // separator
			totalCalories = append(totalCalories, currentElfCalories)
			currentElfCalories = 0

			continue
		}

		num := util.MustParseInt(l)

		currentElfCalories += num
	}

	if currentElfCalories != 0 {
		totalCalories = append(totalCalories, currentElfCalories)
	}

	sort.Ints(totalCalories)

	return totalCalories
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines))
	fmt.Println("Result #2: ", Solution2(lines))
}
