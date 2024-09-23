package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type point models.Point2D

func (p point) String() string {
	return "{" + strconv.Itoa(p.X) + "; " + strconv.Itoa(p.Y) + "}"
}

func (p point) Add(a point) point {
	return point{
		X: p.X + a.X,
		Y: p.Y + a.Y,
	}
}

func (p point) Normalize() (point, int) {
	if p.X != 0 {
		return point{
			X: p.X / utils.Abs(p.X),
		}, utils.Abs(p.X)
	}

	if p.Y != 0 {
		return point{
			Y: p.Y / utils.Abs(p.Y),
		}, utils.Abs(p.Y)
	}

	return point{}, 0
}

type Day struct {
	first, second []point
	origin        point
	field         [][]int
	print         models.PrintFunc
}

func (d *Day) Solve1(input []string, debug bool) string {
	d.Setup(input)
	print := utils.DebugPrint(debug)
	d.print = print

	print("Origin %s. Field size {%d; %d}\n", d.origin.String(), len(d.field), len(d.field[0]))
	res := d.Walk(false)

	return strconv.Itoa(res)
}

func (d *Day) Solve2(input []string, debug bool) string {
	d.Setup(input)
	print := utils.DebugPrint(debug)
	d.print = print

	print("Origin %s. Field size {%d; %d}\n", d.origin.String(), len(d.field), len(d.field[0]))
	res := d.Walk(true)

	return strconv.Itoa(res)
}

func (d *Day) Walk(d2 bool) int {
	// Set minimum as max distance from one corner to another
	var res int
	if d2 {
		res = len(d.field) * len(d.field[0])
	} else {
		res = len(d.field) + len(d.field[0])
	}

	cur := d.origin
	i := 0
	for _, p := range d.first {
		d.print("Step: %s\n", p.String())
		vec, count := p.Normalize()
		for j := 0; j < count; j++ {
			i++
			cur = cur.Add(vec)
			if d.field[cur.Y][cur.X] == 0 {
				d.field[cur.Y][cur.X] = i
			}
		}
	}

	cur = d.origin
	i = 0
	for _, p := range d.second {
		d.print("Step: %s\n", p.String())
		vec, count := p.Normalize()
		for j := 0; j < count; j++ {
			i++
			cur = cur.Add(vec)
			if d.field[cur.Y][cur.X] == 0 {
				continue
			}

			d.print("New intersection at: %s\n", cur.String())
			if d2 {
				res = min(res, i+d.field[cur.Y][cur.X])
			} else {
				res = min(res, utils.ManhDist2D(models.Point2D(d.origin), models.Point2D(cur)))
			}
		}
	}
	return res
}

func (d *Day) Setup(input []string) {
	input = utils.FilterEmptyLines(input)
	if len(input) < 2 {
		fmt.Println("Not enough ropes")
		os.Exit(1)
	}

	d.first = parsePoints(input[0])
	d.second = parsePoints(input[1])

	minF, maxF := calcMinMax(d.first)
	minS, maxS := calcMinMax(d.second)

	d.origin = point{
		X: min(min(minF.X, minS.X), 0),
		Y: min(min(minF.Y, minS.Y), 0),
	}

	lines := max(maxF.Y, maxS.Y) - d.origin.Y + 1
	cols := max(maxF.X, maxS.X) - d.origin.X + 1

	d.field = make([][]int, lines)
	for i := 0; i < lines; i++ {
		d.field[i] = make([]int, cols)
	}

	if d.origin.X < 0 {
		d.origin.X = -d.origin.X
	}

	if d.origin.Y < 0 {
		d.origin.Y = -d.origin.Y
	}
}

func calcMinMax(ps []point) (minP, maxP point) {
	cur := point{}
	for _, p := range ps {
		cur.X += p.X
		cur.Y += p.Y

		if cur.X < minP.X {
			minP.X = cur.X
		}
		if cur.Y < minP.Y {
			minP.Y = cur.Y
		}

		if cur.X > maxP.X {
			maxP.X = cur.X
		}

		if cur.Y > maxP.Y {
			maxP.Y = cur.Y
		}
	}
	return
}

func parsePoints(line string) []point {
	inps := utils.FilterEmptyLines(strings.Split(line, ","))
	res := make([]point, len(inps))
	for i, v := range inps {
		v = strings.TrimSpace(v)
		num, _ := strconv.Atoi(v[1:])
		p := point{}
		switch v[0] {
		case 'L':
			p = point{X: -num}
		case 'R':
			p = point{X: num}
		case 'U':
			p = point{Y: -num}
		case 'D':
			p = point{Y: num}
		}
		res[i] = p
	}
	return res
}
