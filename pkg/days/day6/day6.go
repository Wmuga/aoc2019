package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

type node struct {
	Name     string
	Children []*node
}

const (
	headNode = "COM"
	youNode  = "YOU"
	sanNode  = "SAN"
)

func (d Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	head := parseInput(print, input)

	print("Parsed links: %d\n", len(input))
	print("Links: %+v\n", head)

	count := recWalk(head, 0)

	return strconv.FormatInt(count, 10)
}

func (d Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	head := parseInput(print, input)

	print("Parsed links: %d\n", len(input))
	print("Links: %+v\n", head)

	you, san := recSearch(print, head, 0)
	// -2 since steps to YOU, SAN should not be counted
	return strconv.FormatInt(you+san-2, 10)
}

func recWalk(head *node, counter int64) int64 {
	var count int64
	for _, child := range head.Children {
		count += recWalk(child, counter+1)
	}

	return counter + count
}

func recSearch(print models.PrintFunc, head *node, counter int64) (you int64, san int64) {
	you, san = -1, -1
	switch head.Name {
	case youNode:
		print("YOU node on depth %d\n", counter)
		return counter, -1
	case sanNode:
		print("SAN node on depth %d\n", counter)
		return -1, counter
	}

	for _, child := range head.Children {
		y1, s1 := recSearch(print, child, counter+1)
		// Only found YOU node
		if y1 != -1 && s1 == -1 {
			you = y1
			continue
		}
		// Only found SAN node
		if y1 == -1 && s1 != -1 {
			san = s1
			continue
		}
		// Found both
		if y1 != -1 && s1 != -1 {
			return y1, s1
		}
	}

	// Found YOU and SAN on different children
	if you > 0 && san > 0 {
		print("Found both on %d. DepthYou: %d. DepthSan: %d\n", counter, you, san)
		// sub current path from distances
		you -= counter
		san -= counter
	}

	return
}

func parseInput(print models.PrintFunc, input []string) *node {
	input = utils.FilterEmptyLines(input)

	nodes := map[string]*node{}
	for _, link := range input {
		l := strings.Split(link, ")")
		if len(l) < 2 {
			fmt.Println(link, "No link")
			os.Exit(1)
		}

		print("Link \"%s\" - \"%s\"\n", l[0], l[1])

		nodeR, ok := nodes[l[1]]
		if !ok {
			nodeR = &node{Name: l[1]}
			nodes[l[1]] = nodeR
		}

		nodeL, ok := nodes[l[0]]
		if !ok {
			nodeL = &node{Name: l[0]}
		}

		nodeL.Children = append(nodeL.Children, nodeR)
		nodes[l[0]] = nodeL
	}
	head, ok := nodes[headNode]
	if !ok {
		fmt.Println("Can't find head node", headNode)
		os.Exit(1)
	}
	return head
}
