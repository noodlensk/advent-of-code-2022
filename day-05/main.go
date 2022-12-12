package main

import (
	"fmt"
	"strings"

	"github.com/noodlensk/advent-of-code-2022/util"
)

type CraneOperation struct {
	Qty  int
	From int
	To   int
}

type Stack struct {
	Crates [][]rune
}

func (s *Stack) ExecuteInstruction(operation CraneOperation) {
	fromColIdx := operation.From - 1
	toColIdx := operation.To - 1

	for i := 0; i < operation.Qty; i++ {
		currIdx := len(s.Crates[fromColIdx]) - 1 // last element in stack

		val := s.Crates[fromColIdx][currIdx]
		s.Crates[fromColIdx] = s.Crates[fromColIdx][:currIdx]
		s.Crates[toColIdx] = append(s.Crates[toColIdx], val)
	}
}

func (s *Stack) ExecuteInstructionV2(operation CraneOperation) {
	fromColIdx := operation.From - 1
	toColIdx := operation.To - 1

	currIdxStart := len(s.Crates[fromColIdx]) - operation.Qty // last n element in stack

	vals := s.Crates[fromColIdx][currIdxStart:]
	s.Crates[fromColIdx] = s.Crates[fromColIdx][:currIdxStart]

	s.Crates[toColIdx] = append(s.Crates[toColIdx], vals...)
}

func (s *Stack) Print() {
	maxHeight := 0

	for _, s := range s.Crates {
		if len(s) > maxHeight {
			maxHeight = len(s)
		}
	}

	for i := maxHeight; i > 0; i-- {
		for _, col := range s.Crates {
			if len(col)-1 < i-1 {
				fmt.Printf("    ")

				continue
			}

			fmt.Printf("[%s] ", string(col[i-1]))
		}

		fmt.Println()
	}

	for i := 1; i <= len(s.Crates); i++ {
		fmt.Printf(" %d  ", i)
	}

	fmt.Println()
}

func Solution(lines []string) string {
	sections := strings.Split(strings.Join(lines, "\n"), "\n\n")

	s := mustParseCrateStacks(sections[0])
	instructions := mustParseCraneInstructions(sections[1])

	for _, in := range instructions {
		s.ExecuteInstruction(in)
	}

	topCrates := ""

	for _, col := range s.Crates {
		topCrates += string(col[len(col)-1])
	}

	return topCrates
}

func Solution2(lines []string) string {
	sections := strings.Split(strings.Join(lines, "\n"), "\n\n")

	s := mustParseCrateStacks(sections[0])
	instructions := mustParseCraneInstructions(sections[1])

	for _, in := range instructions {
		s.ExecuteInstructionV2(in)
	}

	topCrates := ""

	for _, col := range s.Crates {
		topCrates += string(col[len(col)-1])
	}

	return topCrates
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines))
	fmt.Println("Result#2: ", Solution2(lines))
}

func mustParseCrateStacks(input string) Stack {
	lines := strings.Split(input, "\n")

	stackSize := len(strings.ReplaceAll(lines[len(lines)-1], " ", "")) // each stack consist of "[R]", and its' 3 symbols per row

	lines = lines[:len(lines)-1] // ignore last row since it's numbers of stacks

	s := Stack{Crates: make([][]rune, stackSize)}

	for i := len(lines) - 1; i >= 0; i-- { // starting from bottom adding items to stacks
		idx := 0

		for j := 1; j < len(lines[i]); j += 4 { // every third symbol
			if lines[i][j] == ' ' {
				idx++

				continue
			}

			s.Crates[idx] = append(s.Crates[idx], rune(lines[i][j]))

			idx++
		}
	}

	return s
}

func mustParseCraneInstructions(input string) []CraneOperation {
	lines := strings.Split(input, "\n")

	var instructions []CraneOperation

	for _, l := range lines {
		parts := strings.Split(l, " ") // Example: "move 1 from 2 to 1"

		instructions = append(instructions, CraneOperation{
			Qty:  util.MustParseInt(parts[1]),
			From: util.MustParseInt(parts[3]),
			To:   util.MustParseInt(parts[5]),
		})
	}

	return instructions
}
