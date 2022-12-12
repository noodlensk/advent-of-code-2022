package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tt := map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	i := 1
	for input, expectedResult := range tt {
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			assert.Equal(t, expectedResult, Solution(input))
		})

		i++
	}
}

func TestSolution2(t *testing.T) {
	tt := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	i := 1
	for input, expectedResult := range tt {
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			assert.Equal(t, expectedResult, Solution2(input))
		})

		i++
	}
}
