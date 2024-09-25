package day7

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
	print := utils.DebugPrint(debug)
	_ = print

	inputs := utils.Permutations([]int64{0, 1, 2, 3, 4}, 0, 4)
	ops := parseInput(input)
	d7vm := vm.New(ops, vm.TypeLogical, false)

	var maxOut int64
PERMLOOP:
	for _, perm := range inputs {
		var outp int64
		for _, setting := range perm {
			d7vm.Reset()
			d7vm.Input([]int64{setting, outp})

			outpAr := d7vm.Run()
			if len(outpAr) == 0 {
				print("Empty output. skip to next permutation")
				continue PERMLOOP
			}
			outp = outpAr[len(outpAr)-1]
		}
		maxOut = max(maxOut, outp)
	}

	return strconv.FormatInt(maxOut, 10)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	_ = print

	inputs := utils.Permutations([]int64{5, 6, 7, 8, 9}, 0, 4)

	ops := parseInput(input)
	vmFunc := func() *vm.VM {
		return vm.New(ops, vm.TypeAwaiter, false)
	}

	var maxOut int64
PERMLOOP2:
	for _, perm := range inputs {
		vms := utils.RepeatFunc(5, vmFunc)
		var outp int64 = 0

		steps := 0
		for vms[4].Status() != vm.StatusHalt {
			for i := range vms {
				if steps == 0 {
					vms[i].Input([]int64{perm[i], outp})
				} else {
					vms[i].Input([]int64{outp})
				}
				vms[i].Continue()

				if vms[i].Status() == vm.StatusError {
					print("VM %d. Error. Skip iteration\n", i)
					continue PERMLOOP2
				}

				outAr := vms[i].Output()
				if len(outAr) == 0 {
					print("VM %d. No output. Skip iteration\n", i)
					continue PERMLOOP2
				}

				outp = outAr[len(outAr)-1]
			}
			steps++
		}

		maxOut = max(maxOut, outp)
	}

	return strconv.FormatInt(maxOut, 10)
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
