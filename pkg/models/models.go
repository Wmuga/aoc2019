package models

type Day interface {
	Solve1(input []string, debug bool) string
	Solve2(input []string, debug bool) string
}

type Test struct {
	Name   string
	Part   int
	Answer string
	Data   []string
}
