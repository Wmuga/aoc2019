package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/days"
	fileparser "github.com/wmuga/aoc2019/pkg/fileParser"
)

const (
	dayNum    = 4
	withPart2 = true
	toTest    = true
)

func getFileNames(day int) (input string, test string) {
	dayStr := strconv.Itoa(day)
	prefix := "inputs/day" + dayStr
	return prefix + "/in.txt", prefix + "/test.txt"
}

func main() {
	day, ok := days.GetDay(dayNum)
	if !ok {
		fmt.Printf("Day %d not found\n", dayNum)
		os.Exit(1)
	}

	inFile, testFile := getFileNames(dayNum)
	inData, err := fileparser.GetInput(inFile)
	if err != nil {
		fmt.Println("Error read input file:", err)
		os.Exit(1)
	}
	testData, err := fileparser.ReadTests(testFile)
	if err != nil {
		fmt.Println("Error read test file:", err)
		os.Exit(1)
	}

	doneTest := true
	if toTest {
		for _, test := range testData {
			var res string
			fmt.Println("Test", test.Name, "for part", test.Part)
			switch test.Part {
			case 1:
				res = day.Solve1(test.Data, true)
			case 2:
				if !withPart2 {
					fmt.Println("Skip part2")
					continue
				}
				res = day.Solve2(test.Data, true)
			default:
				fmt.Println("Skip unknown part", test.Part)
				continue
			}
			if test.Answer == "" {
				fmt.Println("Output:\n", res)
				continue
			}

			out := "[OK]"
			if test.Answer != res {
				out = "[NO]"
				doneTest = false
			}

			fmt.Printf("%s expected: %s, result: %s\n", out, test.Answer, res)
		}
	}

	if !doneTest {
		os.Exit(1)
	}

	fmt.Println("\nAnswers:")
	fmt.Println("Part 1:", day.Solve1(inData, false))

	if withPart2 {
		fmt.Println("Part 2:", day.Solve2(inData, false))
	}
}
