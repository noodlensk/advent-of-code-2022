package main

import (
	"github.com/noodlensk/advent-of-code-2022/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	lines := util.MustReadTestInput()

	res := Solution(lines)
	assert.Equal(t, 95437, res)
}

func TestSolution2(t *testing.T) {
	lines := util.MustReadTestInput()

	res := Solution2(lines)
	assert.Equal(t, 24933642, res)
}