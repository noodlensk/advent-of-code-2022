package main

import (
	"fmt"

	"github.com/noodlensk/advent-of-code-2022/util"
)

type TreeMap struct{ Map [][]int }

func (g TreeMap) MaxX() int { return len(g.Map[0]) - 1 }
func (g TreeMap) MaxY() int { return len(g.Map) - 1 }
func (g TreeMap) Print() {
	for _, row := range g.Map {
		for _, v := range row {
			fmt.Printf("%d", v)
		}

		fmt.Println()
	}
}

// is visible if all points in between are less than dest point
func (g TreeMap) isVisibleFromPoint(xp, yp, x, y int) bool {
	currHeight := g.Map[xp][yp]

	// detect direction
	if xp == x {
		if yp < y {
			for i := yp + 1; i <= y; i++ {
				if g.Map[xp][i] >= currHeight {
					return false
				}
			}
		} else {
			for i := y; i < yp; i++ {
				if g.Map[xp][i] >= currHeight {
					return false
				}
			}
		}
	} else {
		if xp < x {
			for i := xp + 1; i <= x; i++ {
				if g.Map[i][yp] >= currHeight {
					return false
				}
			}
		} else {
			for i := x; i < xp; i++ {
				if g.Map[i][yp] >= currHeight {
					return false
				}
			}
		}
	}

	return true
}

func (g TreeMap) numberOfTreesVisibleFromPoint(xp, yp, x, y int) int {
	currHeight := g.Map[xp][yp]

	cnt := 0
	// detect direction
	if xp == x {
		if yp < y {
			for i := yp + 1; i <= y; i++ {
				cnt++

				if g.Map[xp][i] >= currHeight {
					return cnt
				}
			}
		} else {
			for i := yp - 1; i >= y; i-- {
				cnt++

				if g.Map[xp][i] >= currHeight {
					return cnt
				}
			}
		}
	} else {
		if xp < x {
			for i := xp + 1; i <= x; i++ {
				cnt++

				if g.Map[i][yp] >= currHeight {
					return cnt
				}
			}
		} else {
			for i := xp - 1; i >= x; i-- {
				cnt++

				if g.Map[i][yp] >= currHeight {
					return cnt
				}
			}
		}
	}

	return cnt
}

func (g TreeMap) IsVisible(x, y int) bool {
	return g.isVisibleFromPoint(x, y, 0, y) ||
		g.isVisibleFromPoint(x, y, x, 0) ||
		g.isVisibleFromPoint(x, y, g.MaxX(), y) ||
		g.isVisibleFromPoint(x, y, x, g.MaxY())
}

func (g TreeMap) TreesVisibilityScore(x, y int) int {
	return g.numberOfTreesVisibleFromPoint(x, y, 0, y) *
		g.numberOfTreesVisibleFromPoint(x, y, x, 0) *
		g.numberOfTreesVisibleFromPoint(x, y, g.MaxX(), y) *
		g.numberOfTreesVisibleFromPoint(x, y, x, g.MaxY())
}

func Solution(lines []string) int {
	g := parseTreeMap(lines)
	visibleCnt := 0

	for i := 1; i < g.MaxX(); i++ {
		for j := 1; j < g.MaxY(); j++ {
			if g.IsVisible(i, j) {
				visibleCnt++
			}
		}
	}

	visibleCnt += (g.MaxX() + g.MaxY()) * 2

	return visibleCnt
}

func Solution2(lines []string) int {
	treeMap := parseTreeMap(lines)
	bestTreeScore := 0

	for i := 1; i < treeMap.MaxX(); i++ {
		for j := 1; j < treeMap.MaxY(); j++ {
			currScore := treeMap.TreesVisibilityScore(i, j)

			if currScore > bestTreeScore {
				bestTreeScore = currScore
			}
		}
	}

	return bestTreeScore
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines))
	fmt.Println("Result #2: ", Solution2(lines))
}

func parseTreeMap(lines []string) TreeMap {
	treeMap := TreeMap{}

	for _, l := range lines {
		var row []int

		for _, c := range l {
			row = append(row, util.MustParseInt(string(c)))
		}

		treeMap.Map = append(treeMap.Map, row)
	}

	return treeMap
}
