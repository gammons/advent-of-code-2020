package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	StepRight = 3
	StepDown  = 1
)

func main() {
	x, y, treeCount := 0, 0, 0
	world := readWorld()

	for y < len(world) {
		if x%(len(world[y])-1) > 0 {
			x = x % (len(world[y]))
		}

		// fmt.Printf("(%d, %d)\n", x, y)
		if world[y][x] == "#" {
			treeCount++
		}

		x += StepRight
		y += StepDown
	}

	fmt.Println("TreeCount = ", treeCount)
}

func readWorld() [][]string {
	data, _ := ioutil.ReadFile("input.txt")
	var input [][]string

	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			input = append(input, strings.Split(line, ""))
		}
	}
	return input
}

func loadExampleWorld() [][]string {
	var ret [][]string

	ret = append(ret, strings.Split("..##.......", ""))
	ret = append(ret, strings.Split("#...#...#..", ""))
	ret = append(ret, strings.Split(".#....#..#.", ""))
	ret = append(ret, strings.Split("..#.#...#.#", ""))
	ret = append(ret, strings.Split(".#...##..#.", ""))
	ret = append(ret, strings.Split("..#.##.....", ""))
	ret = append(ret, strings.Split(".#.#.#....#", ""))
	ret = append(ret, strings.Split(".#........#", ""))
	ret = append(ret, strings.Split("#.##...#...", ""))
	ret = append(ret, strings.Split("#...##....#", ""))
	ret = append(ret, strings.Split(".#..#...#.#", ""))
	return ret
}
