package day5

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2019/pkg/vm"
)

const opcodeInp = 3

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	ops := parseInput(input)
	d5vm := vm.New(ops, vm.TypeInOut, debug)

	inpLen := utils.CountFunc(ops, func(item int64) bool {
		return item%100 == opcodeInp
	})

	inp := slices.Repeat([]int64{1}, inpLen)
	d5vm.Input(inp)
	d5vm.Run()

	res := d5vm.Output()

	return utils.JoinInt64(res, ",")
}

func (Day) Solve2(input []string, debug bool) string {
	ops := parseInput(input)
	d5vm := vm.New(ops, vm.TypeLogical, debug)

	inpLen := utils.CountFunc(ops, func(item int64) bool {
		return item%100 == opcodeInp
	})

	inp := slices.Repeat([]int64{5}, inpLen)
	d5vm.Input(inp)
	d5vm.Run()

	res := d5vm.Output()

	return utils.JoinInt64(res, ",")
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
