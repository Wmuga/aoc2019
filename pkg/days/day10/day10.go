package day10

import (
	"math"
	"sort"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/set"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

type point models.Point2D

func (p point) Add(p2 point) point {
	return point{p.X + p2.X, p.Y + p2.Y}
}

func (p point) Mul(n int) point {
	return point{p.X * n, p.Y * n}
}

func (p point) String() string {
	return strconv.Itoa(p.X) + ";" + strconv.Itoa(p.Y)
}

func (p point) Hash() string {
	return p.String()
}

const rAsteroid = '#'

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	points, field := parseInput(input)
	print("Field: %d:%d. Asteroids: %d\n", len(field[0]), len(field), len(points))
	p, count := findMax(points, field)
	print("Max point: %d:%d with count: %d\n", p.X, p.Y, count)
	return strconv.Itoa(count)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	points, field := parseInput(input)
	print("Field: %d:%d. Asteroids: %d\n", len(field[0]), len(field), len(points))
	p, count := findMax(points, field)
	print("Max point: %d:%d with count: %d\n", p.X, p.Y, count)

	vecs := getVectors(p, points)
	sort.SliceStable(vecs, func(i, j int) bool {
		return angle(vecs[i]) < angle(vecs[j])
	})

	var (
		last    point
		i       int
		counter int
	)
	for counter < 200 {
		vec := vecs[i%len(vecs)]
		last = p.Add(vec.Mul(i/len(vecs) + 1))
		i++
		for last.X >= 0 && last.X < len(field[0]) && last.Y >= 0 && last.Y < len(field) {
			if field[last.Y][last.X] != rAsteroid {
				last = last.Add(vec)
				continue
			}

			counter++
			break

		}
	}

	return strconv.Itoa(last.X*100 + last.Y)
}

func findMax(points []point, field [][]rune) (p point, maxCount int) {
	for _, cur := range points {
		vectors := getVectors(cur, points)
		count := len(vectors)

		if count > maxCount {
			maxCount = count
			p = cur
		}
	}
	return
}

func getVectors(cur point, points []point) []point {
	vecs := set.NewHasherSet[point]()

	for _, p := range points {
		x := p.X - cur.X
		y := p.Y - cur.Y

		if x == 0 && y == 0 {
			continue
		}

		if x == 0 {
			vecs.Upsert(point{0, y / utils.Abs(y)})
			continue
		}

		if y == 0 {
			vecs.Upsert(point{x / utils.Abs(x), 0})
			continue
		}

		gcd := utils.GCD(utils.Abs(x), utils.Abs(y))
		vecs.Upsert(point{x / gcd, y / gcd})
	}

	return utils.ToSlice(vecs.Iterator())
}

func parseInput(input []string) (points []point, field [][]rune) {
	input = utils.FilterEmptyLines(input)
	field = make([][]rune, len(input))
	for i := range field {
		for j, c := range []rune(input[i]) {
			if c == rAsteroid {
				points = append(points, point{j, i})
			}
		}
		field[i] = []rune(input[i])
	}

	return
}

// angle returns angle of point to point{0,1}
func angle(p point) float64 {
	angle := math.Acos(-float64(p.Y) / math.Sqrt(float64(p.X*p.X+p.Y*p.Y)))
	if p.X >= 0 {
		return angle
	}
	return 2*math.Pi - angle
}
