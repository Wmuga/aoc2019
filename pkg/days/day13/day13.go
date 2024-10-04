package day13

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
	ops := parseInput(input)
	d13vm := vm.New(ops, vm.TypeMemory, debug)
	out := d13vm.Run()
	count := 0
	for i := 2; i < len(out); i += 3 {
		tileType := out[i]
		if tileType == 2 {
			count++
		}
	}
	return strconv.Itoa(count)
}

func (Day) Solve2(input []string, debug bool) string {
	ops := parseInput(input)
	d13vm := vm.New(ops, vm.TypeMemory, debug)
	d13vm.SetAt(0, 2)

	var (
		xPlayer int64
		xBall   int64
		score   int64
	)

	for {
		out := d13vm.Continue()
		for i := 2; i < len(out); i += 3 {
			x := out[i-2]
			y := out[i-1]
			tileType := out[i]

			if x == -1 && y == 0 {
				score = tileType
				continue
			}

			switch tileType {
			case 3:
				xPlayer = x
			case 4:
				xBall = x
			}
		}

		if d13vm.Status() != vm.StatusAwaitInput {
			break
		}

		inp := xBall - xPlayer
		if inp != 0 {
			inp = inp / utils.Abs(inp)
		}

		d13vm.Input([]int64{inp})
	}
	return strconv.FormatInt(score, 10)
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
