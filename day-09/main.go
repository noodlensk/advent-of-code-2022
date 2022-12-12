package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Steps     int
}

type Grid struct {
	state [][]int // 0 - unvisited, 1 visited
	head  Point
	tail  Point
}

func (g *Grid) Print() {
	for i := 0; i < len(g.state[0]); i++ {
		for j := 0; j < len(g.state[0]); j++ {
			fmt.Printf(".")
		}

		fmt.Println()
	}
}

func NewGrid(x, y int) Grid {
	g := Grid{}
	for i := 0; i < y; i++ {
		g.state = append(g.state, make([]int, x))
	}

	return g
}

type Point struct {
	X int
	Y int
}

func Solution(lines []string) int  { return 0 }
func Solution2(lines []string) int { return 0 }

func main() {
	fileData, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	instructions := parseInstructions(string(fileData))

	var (
		minY, maxY, currY int
		minX, maxX, currX int
	)

	for _, i := range instructions {
		if i.Direction == "U" {
			currY += i.Steps
		}

		if i.Direction == "D" {
			currY -= i.Steps
		}

		if currY > maxY {
			maxY = currY
		}

		if currY < minY {
			minY = currY
		}

		if i.Direction == "R" {
			currX += i.Steps
		}

		if i.Direction == "L" {
			currX -= i.Steps
		}

		if currX > maxX {
			maxX = currX
		}

		if currX < minX {
			minX = currX
		}
	}

	fmt.Printf("Size: [%d, %d], init point [%d, %d]\n", maxX-minX, maxY-minY, 0-minX, 0-minY)

	g := NewGrid(maxX-minX, maxY-minY)

	g.Print()
}

func parseInstructions(input string) []Instruction {
	lines := strings.Split(input, "\n")

	var res []Instruction

	for _, l := range lines {
		i := Instruction{Direction: string(l[0])}

		steps, err := strconv.ParseInt(l[2:], 10, 32) // first two symbols is direction and space
		if err != nil {
			panic(err)
		}

		i.Steps = int(steps)

		res = append(res, i)
	}

	return res
}
