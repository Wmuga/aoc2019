package models

type Day interface {
	Solve1(input []string) string
	Solve2(input []string) string
}

type Test struct {
	Name   string
	Part   int
	Answer string
	Data   []string
}
