package main

import (
	"fmt"
	"strings"

	"github.com/noodlensk/advent-of-code-2022/util"
)

func Solution(lines []string) int {
	itemsPriority := 0

	for _, rucksack := range lines {
		commonItem := mustFindCommonItemInCompartments(rucksack)
		if commonItem != nil {
			itemsPriority += itemPriority(*commonItem)
		}
	}

	return itemsPriority
}

func Solution2(lines []string) int {
	itemsPriority := 0

	for i := 0; i < len(lines); i += 3 {
		commonItem := findCommonItem(lines[i], lines[i+1], lines[i+2])
		if commonItem != nil {
			itemsPriority += itemPriority(*commonItem)
		}
	}

	return itemsPriority
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Priority sum: ", Solution(lines))
	fmt.Println("Priority sum #2: ", Solution2(lines))
}

func mustFindCommonItemInCompartments(rucksack string) *rune {
	if len(rucksack)%2 != 0 {
		panic("need to have even rucksack length")
	}

	compartmentSize := len(rucksack) / 2

	itemsTracker := map[rune]int{} // -1 for first, 1 for second compartment

	for i := 0; i < len(rucksack); i++ {
		item := rune(rucksack[i])

		if itemsTracker[item] < 0 && i >= compartmentSize {
			return &item
		}

		if i < compartmentSize {
			itemsTracker[item] = -1
		} else {
			itemsTracker[item] = 1
		}
	}

	return nil
}

func itemPriority(r rune) int {
	if strings.ToUpper(string(r)) == string(r) {
		return int(r-64) + 26
	}

	return int(r - 96)
}

func findCommonItem(r1, r2, r3 string) *rune {
	itemTracker := map[rune]map[int]bool{} // map[item][rucksackIdx]true

	for rucksackIdx, rucksack := range []string{r1, r2, r3} {
		for _, item := range rucksack {
			if itemTracker[item] == nil {
				itemTracker[item] = map[int]bool{}
			}

			itemTracker[item][rucksackIdx] = true
		}
	}

	for r, v := range itemTracker {
		if len(v) == 3 { // in all 3 rucksacks
			return &r
		}
	}

	return nil
}
