package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2019/pkg/vm"
)

type Day struct{}

type point models.Point2D

func (p point) Add(p2 point) point { return point{p.X + p2.X, p.Y + p2.Y} }

var (
	vecL = point{-1, 0}
	vecR = point{1, 0}
	vecU = point{0, -1}
	vecD = point{0, 1}
)

func (Day) Solve1(input []string, debug bool) string {
	c, _ := solve(input, debug, false)
	return strconv.Itoa(c)
}

func (Day) Solve2(input []string, debug bool) string {
	_, res := solve(input, debug, true)
	return res
}

func solve(input []string, debug bool, startWhite bool) (count int, out string) {
	cur := point{}
	vec := vecU
	field := map[point]bool{
		cur: startWhite,
	}

	ops := parseInput(input)
	d11vm := vm.New(ops, vm.TypeMemory, debug)

	var inp int64
	if startWhite {
		inp = 1
	}
	d11vm.Input([]int64{inp})

	for {
		outAr := d11vm.Continue()

		if d11vm.Status() != vm.StatusAwaitInput {
			break
		}
		if len(outAr) < 2 {
			fmt.Println("Insufficient output")
			return
		}

		field[cur] = outAr[0] == 1
		switch vec {
		case vecU:
			vec = utils.OneOf(outAr[1] == 0, vecL, vecR)
		case vecL:
			vec = utils.OneOf(outAr[1] == 0, vecD, vecU)
		case vecD:
			vec = utils.OneOf(outAr[1] == 0, vecR, vecL)
		case vecR:
			vec = utils.OneOf(outAr[1] == 0, vecU, vecD)
		}

		cur = cur.Add(vec)

		inp = 0
		if field[cur] {
			inp = 1
		}

		d11vm.Input([]int64{inp})

	}

	var minX, maxX, minY, maxY int

	for k := range field {
		minX = min(k.X, minX)
		maxX = max(k.X, maxX)
		minY = min(k.Y, minY)
		maxY = max(k.Y, maxY)
	}

	pic := make([][]rune, maxY-minY+1)
	for i := range pic {
		pic[i] = []rune(strings.Repeat(" ", maxX-minX+1))
	}

	for k, v := range field {
		if v {
			x := k.X - minX
			y := k.Y - minY
			pic[y][x] = '#'
		}
		count++
	}

	for _, line := range pic {
		out = out + "\n" + string(line)
	}

	return
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
