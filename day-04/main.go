package main

import (
	"fmt"
	"strings"

	"github.com/noodlensk/advent-of-code-2022/util"
)

type Range struct {
	From int
	To   int
}

func (r Range) Contains(d Range) bool {
	if r.From >= d.From && r.To <= d.To {
		return true
	}

	return false
}

func (r Range) Overlaps(d Range) bool {
	if r.From <= d.To && r.To >= d.From {
		return true
	}

	return false
}

func Solution(lines []string) int {
	count := 0

	for _, l := range lines {
		a, b := mustParseRanges(l)

		if a.Contains(b) || b.Contains(a) {
			count++
		}
	}

	return count
}

func Solution2(lines []string) int {
	count := 0

	for _, l := range lines {
		a, b := mustParseRanges(l)

		if a.Overlaps(b) || b.Overlaps(a) {
			count++
		}
	}

	return count
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines))
	fmt.Println("Result #2: ", Solution2(lines))
}

func mustParseRanges(input string) (a, b Range) {
	parts := strings.Split(input, ",")

	partsA := strings.Split(parts[0], "-")
	partsB := strings.Split(parts[1], "-")

	a.From = util.MustParseInt(partsA[0])
	a.To = util.MustParseInt(partsA[1])

	b.From = util.MustParseInt(partsB[0])
	b.To = util.MustParseInt(partsB[1])

	return
}
