package day12

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

type point models.Point3D

type pair models.Point2D

func (p point) Add(p2 point) point {
	return point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}

func (p point) Inv() point {
	return point{-p.X, -p.Y, -p.Z}
}

func (p point) Sub(p2 point) point {
	return p.Add(point{-p2.X, -p2.Y, -p2.Z})
}

func (p point) Normalize() point {
	x := 0
	if p.X > 0 {
		x = 1
	}
	if p.X < 0 {
		x = -1
	}
	y := 0
	if p.Y > 0 {
		y = 1
	}
	if p.Y < 0 {
		y = -1
	}
	z := 0
	if p.Z > 0 {
		z = 1
	}
	if p.Z < 0 {
		z = -1
	}

	return point{x, y, z}
}

func (p point) String() string {
	return fmt.Sprintf("%d;%d;%d", p.X, p.Y, p.Z)
}

type moon struct {
	pos point
	vel point
}

var reDigit = regexp.MustCompile(`[+-]?\d+`)

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	moons := parseInput(input)
	print("Parsed moons: %+v\n", moons)
	for i := 0; i < 1000; i++ {
		moons = step(moons)
	}

	var sum int64
	for _, m := range moons {
		p := utils.Abs(m.pos.X) + utils.Abs(m.pos.Y) + utils.Abs(m.pos.Z)
		k := utils.Abs(m.vel.X) + utils.Abs(m.vel.Y) + utils.Abs(m.vel.Z)
		sum += int64(p * k)
	}

	return strconv.FormatInt(sum, 10)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	moons := parseInput(input)
	print("Parsed moons: %+v\n", moons)
	var (
		steps   = [3]int64{0, 0, 0}
		counter int64

		states = [3][]pair{calcState(moons, 0), calcState(moons, 1), calcState(moons, 2)}
	)

	for steps[0] == 0 || steps[1] == 0 || steps[2] == 0 {
		counter++
		moons = step(moons)

		for i, state := range states {
			if steps[i] != 0 {
				continue
			}

			curState := calcState(moons, i)
			if stateEQ(curState, state) {
				steps[i] = counter
			}
		}
	}

	res := utils.LCM(utils.LCM(steps[0], steps[1]), steps[2])

	return strconv.FormatInt(res, 10)
}

func calcState(moons []moon, coord int) []pair {
	var f func(point) int
	switch coord {
	case 0:
		f = func(p point) int { return p.X }
	case 1:
		f = func(p point) int { return p.Y }
	case 2:
		f = func(p point) int { return p.Z }
	default:
		return nil
	}

	res := make([]pair, len(moons))
	for i := range moons {
		res[i] = pair{f(moons[i].pos), f(moons[i].vel)}
	}

	return res
}

func stateEQ(s1 []pair, s2 []pair) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// step simulates 1 step in simulation. mutates original array
func step(state []moon) []moon {
	// calculate velocity
	for i := 0; i < len(state)-1; i++ {
		m1 := state[i]
		for j := i + 1; j < len(state); j++ {
			m2 := state[j]
			vel := m2.pos.Sub(m1.pos).Normalize()
			m1.vel = m1.vel.Add(vel)
			m2.vel = m2.vel.Add(vel.Inv())
			state[j] = m2
		}
		state[i] = m1
	}
	// apply velocity
	for i := 0; i < len(state); i++ {
		m := state[i]
		m.pos = m.pos.Add(m.vel)
		state[i] = m
	}

	return state
}

func parseInput(input []string) []moon {
	input = utils.FilterEmptyLines(input)
	moons := make([]moon, len(input))
	for i := range input {
		coords := reDigit.FindAllString(input[i], -1)
		if len(coords) != 3 {
			fmt.Println("Wrong number count per line", coords)
			os.Exit(1)
		}
		coordsInt, err := utils.ParseIntLines(coords)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		moons[i] = moon{pos: point{coordsInt[0], coordsInt[1], coordsInt[2]}}
	}
	return moons
}
