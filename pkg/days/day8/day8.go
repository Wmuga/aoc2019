package day8

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

func isNumber(b byte) func(byte) bool {
	return func(bcmp byte) bool {
		return b == bcmp
	}
}

func countNums(b byte, layer [][]byte) int {
	countFunc := func(b1 byte) bool { return b == b1 }
	count := 0
	for i := 0; i < len(layer); i++ {
		count += utils.CountFunc(layer[i], countFunc)
	}
	return count
}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	img := parseInput(input)
	print("Image: %d:%d:%d\n", len(img), len(img[0][0]), len(img[0]))

	minLayer := 0
	minCount := countNums(0, img[0])

	for l := 0; l < len(img); l++ {
		count := countNums(0, img[l])
		if count >= minCount {
			continue
		}
		minLayer = l
		minCount = count
	}
	print("Min layer: %d; min count: %d\n", minLayer, minCount)
	ones := countNums(1, img[minLayer])
	twos := countNums(2, img[minLayer])
	print("Ones: %d Twos: %d\n", ones, twos)
	res := ones * twos

	return strconv.Itoa(res)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	img := parseInput(input)
	print("Image: %d:%d:%d\n", len(img), len(img[0][0]), len(img[0]))

	fullImg := make([][]rune, len(img[0]))
	for i := range fullImg {
		fullImg[i] = make([]rune, len(img[0][0]))
		for j := range fullImg[i] {
			fullImg[i][j] = ' '
		}
	}

	for y := range fullImg {
		for x := range fullImg[y] {
		layerloop:
			for l := range img {
				switch img[l][y][x] {
				case 1:
					fullImg[y][x] = '#'
				case 0:
					break layerloop
				}
			}
		}
	}

	res := ""
	for _, line := range fullImg {
		res += "\n" + string(line)
	}

	return res
}

func parseInput(input []string) [][][]byte {
	input = utils.FilterEmptyLines(input)
	if len(input) < 3 {
		fmt.Println("Wrong input")
		os.Exit(1)
	}

	width := utils.Must(utils.ParseInt(input[0]))
	height := utils.Must(utils.ParseInt(input[1]))
	layers := len(input[2]) / width / height

	res := make([][][]byte, layers)
	for l := 0; l < layers; l++ {
		layer := make([][]byte, height)
		for i := 0; i < height; i++ {
			line := make([]byte, width)
			for j := 0; j < width; j++ {
				line[j] = input[2][l*width*height+i*width+j] - '0'
			}
			layer[i] = line
		}
		res[l] = layer
	}

	return res
}
