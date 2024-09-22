package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2019/pkg/vm"
)

type Day struct{}

func (d Day) Solve1(input []string, debug bool) string {
	ops := parseInput(input)
	d2VM := vm.New(ops, vm.TypeSimple, debug)

	// only on actual data
	if !debug {
		d2VM.SetAt(1, 12)
		d2VM.SetAt(2, 2)
	}

	// steps vm
	for d2VM.Next() {
	}

	res, _ := d2VM.GetAt(0)

	return strconv.FormatInt(res, 10)
}

func (d Day) Solve2(input []string, debug bool) string {
	// no tests for part2
	if debug {
		return ""
	}

	ops := parseInput(input)
	d2VM := vm.New(ops, vm.TypeSimple, debug)

	var (
		res  int64
		noun int64
		verb int64
	)

	for noun != 99 || verb != 99 {
		d2VM.Reset()
		d2VM.SetAt(1, noun)
		d2VM.SetAt(2, verb)
		for d2VM.Next() {
		}
		res, _ = d2VM.GetAt(0)

		if res == 19690720 {
			break
		}

		verb++
		if verb == 100 {
			noun++
			verb = 0
		}
	}

	return strconv.FormatInt(noun*100+verb, 10)
}

func parseInput(input []string) []int64 {
	ops := strings.Split(input[0], ",")
	res := make([]int64, len(ops))
	i := 0
	for num, err := range utils.ParseIntRange(ops) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		res[i] = int64(num)
		i++
	}
	return res
}
