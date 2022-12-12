package util

import "strconv"

func MustParseInt(input string) int {
	num, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(num)
}
