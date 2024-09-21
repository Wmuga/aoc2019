package day1

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

func (d Day) Solve1(input []string, debug bool) string {
	input = utils.FilterEmptyLines(input)
	var acc int64

	for num, err := range utils.ParseIntRange(input) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		acc += calc(int64(num))
	}

	return strconv.FormatInt(acc, 10)
}

func (d Day) Solve2(input []string, debug bool) string {
	input = utils.FilterEmptyLines(input)
	var acc int64

	for num, err := range utils.ParseIntRange(input) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val := calc(int64(num))
		for val > 0 {
			acc += val
			val = calc(val)
		}
	}

	return strconv.FormatInt(acc, 10)
}

func calc(a int64) int64 {
	val := a/3 - 2
	if val < 0 {
		return 0
	}

	return val
}
