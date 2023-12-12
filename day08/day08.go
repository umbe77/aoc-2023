package day08

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	in := make([]string, 0)
	utils.ReadFile("day08/input.txt", func(line string) {
		in = append(in, line)
	})

	nm := parse(in)

	fmt.Printf("Part 1: %d\n", part1(nm))
	fmt.Printf("Part 2: %d\n", part2(nm))
}

type Node struct {
	end   string
	left  string
	right string
}

type NetMap struct {
	nodes map[string]Node
	path  string
	start string
}

func parse(input []string) NetMap {
	m := NetMap{
		nodes: make(map[string]Node),
		path:  input[0],
	}

	for i := 2; i < len(input); i++ {
		r := input[i]
		nodepart := strings.Split(r, "=")
		node := strings.TrimSpace(nodepart[0])
		childs := strings.Split(strings.Trim(strings.TrimSpace(nodepart[1]), "()"), ",")

		n := Node{
			end:   string(node[2]),
			left:  strings.Trim(childs[0], " "),
			right: strings.Trim(childs[1], " "),
		}
		if i == 2 {
			m.start = node
		}
		m.nodes[node] = n
	}

	return m
}

// func getNext(nodes map[string]Node, current Node, dir string)

func part1(m NetMap) int {
	moves := 0
	for k := range m.nodes {
		if k == "AAA" {
			m.start = k
			break
		}
	}
	current := m.start
	// fmt.Printf("moves: %d -- %+v\n", moves, m)
	for {
		// fmt.Printf("moves: %d -- current: %s\n", moves, current)
		if current == "ZZZ" {
			break
		}
		for _, v := range m.path {
			// fmt.Printf("movepath: %s -- ", string(v))
			moves = moves + 1
			if v == 'L' {
				current = m.nodes[current].left
				// fmt.Printf("moves: %d, %s -- current: %s\n", moves, string(v), current)
				continue
			}
			current = m.nodes[current].right
			// fmt.Printf("moves: %d, %s -- current: %s\n", moves, string(v), current)
		}
	}
	return moves
}

func part2(m NetMap) int {
	starts := make([]string, 0)
	for k := range m.nodes {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

    // fmt.Printf("nodemoves: %v\n", starts)
	nodeMoves := make([]int, len(starts))
	for i, n := range starts {
		moves := 0
        current := n
		for {
			// fmt.Printf("moves: %d -- current: %s\n", moves, current)
			if current[2] == 'Z' {
				break
			}
			for _, v := range m.path {
				// fmt.Printf("movepath: %s -- ", string(v))
				moves = moves + 1
				if v == 'L' {
					current = m.nodes[current].left
					// fmt.Printf("moves: %d, %s -- current: %s\n", moves, string(v), current)
					continue
				}
				current = m.nodes[current].right
				// fmt.Printf("moves: %d, %s -- current: %s\n", moves, string(v), current)
			}
		}
        nodeMoves[i] = moves
	}

    // fmt.Printf("nodemoves: %v\n", nodeMoves)
    // fmt.Printf("2 3 -- %d %d\n", utils.Gcd(2, 3), utils.MustLcm([]int{2, 3, 5, 7}))
    return utils.MustLcm(nodeMoves)
}
