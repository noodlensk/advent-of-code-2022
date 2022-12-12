package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/noodlensk/advent-of-code-2022/util"
)

func TestSolution(t *testing.T) {
	lines := util.MustReadTestInput()

	res := Solution(lines)
	assert.Equal(t, 24000, res)
}

func TestSolution2(t *testing.T) {
	lines := util.MustReadTestInput()

	res := Solution2(lines)
	assert.Equal(t, 45000, res)
}
