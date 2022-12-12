package main

import (
	"fmt"

	"github.com/noodlensk/advent-of-code-2022/util"
)

type GameAction int

const (
	GameActionRock = iota + 1
	GameActionPaper
	GameActionScissors
)

type GameResult int

const (
	GameResultWin = iota + 1
	GameResultDraw
	GameResultLoose
)

func Solution(lines []string) int {
	score := 0

	for _, l := range lines {
		actionA := mustParseGameAction(string(l[0]))
		actionB := mustParseGameAction(string(l[2]))

		gameResult := ComputeGameResult(actionA, actionB)

		score += totalScore(actionA, actionB, gameResult)
	}

	return score
}

func Solution2(lines []string) int {
	score := 0

	for _, l := range lines {
		actionA := mustParseGameAction(string(l[0]))
		desiredResult := mustParseGameResult(string(l[2]))
		actionB := mustRecommendGameAction(actionA, desiredResult)

		gameResult := ComputeGameResult(actionA, actionB)

		score += totalScore(actionA, actionB, gameResult)
	}

	return score
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Total score: ", Solution(lines))
	fmt.Println("Total score #2: ", Solution2(lines))
}

func totalScore(_, b GameAction, gameResult GameResult) int {
	score := 0

	switch gameResult {
	case GameResultWin:
		score += 6
	case GameResultDraw:
		score += 3
	case GameResultLoose:
		// nothing
	}

	switch b {
	case GameActionRock:
		score += 1 //nolint:revive
	case GameActionPaper:
		score += 2
	case GameActionScissors:
		score += 3
	}

	return score
}

func ComputeGameResult(a, b GameAction) GameResult {
	switch {
	case a == b:
		return GameResultDraw
	case a == GameActionRock && b == GameActionPaper:
		return GameResultWin
	case a == GameActionRock && b == GameActionScissors:
		return GameResultLoose
	case a == GameActionPaper && b == GameActionRock:
		return GameResultLoose
	case a == GameActionPaper && b == GameActionScissors:
		return GameResultWin
	case a == GameActionScissors && b == GameActionRock:
		return GameResultWin
	case a == GameActionScissors && b == GameActionPaper:
		return GameResultLoose
	}

	return GameResultDraw
}

func mustRecommendGameAction(a GameAction, r GameResult) GameAction {
	switch r {
	case GameResultDraw:
		return a
	case GameResultWin:
		switch a {
		case GameActionRock:
			return GameActionPaper
		case GameActionPaper:
			return GameActionScissors
		case GameActionScissors:
			return GameActionRock
		}
	case GameResultLoose:
		switch a {
		case GameActionRock:
			return GameActionScissors
		case GameActionPaper:
			return GameActionRock
		case GameActionScissors:
			return GameActionPaper
		}
	default:
		panic("unknown")
	}

	return 0
}

func mustParseGameAction(s string) GameAction {
	switch s {
	case "A":
		return GameActionRock
	case "B":
		return GameActionPaper
	case "C":
		return GameActionScissors
	case "X":
		return GameActionRock
	case "Y":
		return GameActionPaper
	case "Z":
		return GameActionScissors
	default:
		panic("unknown")
	}
}

func mustParseGameResult(s string) GameResult {
	switch s {
	case "X":
		return GameResultLoose
	case "Y":
		return GameResultDraw
	case "Z":
		return GameResultWin
	default:
		panic("unknown")
	}
}
