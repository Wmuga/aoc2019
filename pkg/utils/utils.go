package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
)

// FilterEmptyLines deletes empty entries in slice
func FilterEmptyLines(lines []string) []string {
	return slices.DeleteFunc(lines, func(s string) bool { return strings.TrimSpace(s) == "" })
}

// ParseInt parses line to int
func ParseInt(line string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(line))
}

// ParseIntLines parses lines to ints
func ParseIntLines(lines []string) ([]int, error) {
	res := make([]int, len(lines))
	for i, line := range lines {
		num, err := ParseInt(line)
		if err != nil {
			return nil, err
		}
		res[i] = num
	}
	return res, nil
}

// ParseIntRange returns range func for parsing lines to int
func ParseIntRange(lines []string) func(yield func(int, error) bool) {
	return func(yield func(int, error) bool) {
		for _, line := range lines {
			num, err := ParseInt(line)
			// error - return error and be gone
			if err != nil {
				yield(0, err)
				return
			}
			// try return num
			if !yield(num, nil) {
				return
			}
		}
	}
}

// ManhDist2D returns manhattan distance between 2 points on plane
func ManhDist2D(a, b models.Point2D) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// DebugPrint prints formated text if debug=true
func DebugPrint(debug bool) models.PrintFunc {
	if !debug {
		return func(string, ...interface{}) {}
	}

	return func(format string, data ...interface{}) {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}

		fmt.Printf(format, data...)
	}
}

func Count[T comparable](s []T, item T) int {
	return CountFunc(s, func(item2 T) bool { return item == item2 })
}

func CountFunc[T comparable](s []T, f func(T) bool) int {
	count := 0
	for i := range s {
		if f(s[i]) {
			count++
		}
	}
	return count
}

func JoinInt64(s []int64, delim string) string {
	s2 := make([]string, len(s))
	for i := range s {
		s2[i] = strconv.FormatInt(s[i], 10)
	}
	return strings.Join(s2, delim)
}

func Permutations[T any](inp []T, start, end int) [][]T {
	var res [][]T

	if start == end {
		return [][]T{inp}
	}

	for i := start; i <= end; i++ {
		newInp := slices.Clone(inp)
		newInp[start], newInp[i] = newInp[i], newInp[start]
		res = append(res, Permutations(newInp, start+1, end)...)
	}

	return res
}

func RepeatFunc[T any](count int, f func() T) []T {
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = f()
	}
	return res
}
