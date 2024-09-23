package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	start, end := parseInput(input)
	print("Range: %d-%d. Part 1\n", start, end)
	count := checkRange(print, start, end, false)

	return strconv.Itoa(count)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	start, end := parseInput(input)
	print("Range: %d-%d. Part 2\n", start, end)
	count := checkRange(print, start, end, true)

	return strconv.Itoa(count)
}

func checkRange(print models.PrintFunc, start, end int, d2 bool) (count int) {
	// count repeating numbers
	var counts = []int{0}
OUTER:
	for i := start; i <= end; i++ {
		// drop last counts. keep memory
		counts = counts[:0]
		nums := toNums(i)
		counts = append(counts, 1)
	INNER:
		for j := 1; j < len(nums); j++ {
			// Condition 1. Non-descreasing digits
			if nums[j-1] > nums[j] {
				continue OUTER
			}

			// Count repeating
			if nums[j-1] == nums[j] {
				counts[len(counts)-1]++
				continue INNER
			}

			counts = append(counts, 1)
		}
		// Condition 2 and 3. Minimum 2 repeating. Only 2 repeating(part 2)
		for _, c := range counts {
			if c == 2 || (!d2 && c > 2) {
				count++
				continue OUTER
			}
		}
	}
	return
}

func parseInput(inp []string) (start, end int) {
	inp = utils.FilterEmptyLines(inp)
	if len(inp) == 0 {
		fmt.Println("No input")
		os.Exit(1)
	}

	inp = utils.FilterEmptyLines(strings.Split(inp[0], "-"))

	if len(inp) < 2 {
		fmt.Println("No input range")
		os.Exit(1)
	}

	start, err := strconv.Atoi(inp[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	end, err = strconv.Atoi(inp[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func toNums(num int) []byte {
	numStr := []byte(strconv.Itoa(num))
	for i := range numStr {
		numStr[i] -= '0'
	}
	return numStr
}
