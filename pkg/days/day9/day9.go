package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2019/pkg/vm"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, 1)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, 2)
}

func solve(input []string, debug bool, in int64) string {
	ops := parseInput(input)
	d9vm := vm.New(ops, vm.TypeMemory, debug)
	d9vm.Input([]int64{in})
	d9vm.Run()
	out := d9vm.Output()
	if len(out) == 0 {
		return ""
	}
	return strconv.FormatInt(out[0], 10)

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
