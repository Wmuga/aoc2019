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
