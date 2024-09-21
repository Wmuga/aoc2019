package utils

import (
	"slices"
	"strconv"
	"strings"
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
