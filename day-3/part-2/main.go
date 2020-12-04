package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	world := loadWorld()
	tc1 := treeCountForSlope(world, 1, 1)
	tc2 := treeCountForSlope(world, 3, 1)
	tc3 := treeCountForSlope(world, 5, 1)
	tc4 := treeCountForSlope(world, 7, 1)
	tc5 := treeCountForSlope(world, 1, 2)

	fmt.Println("product = ", tc1*tc2*tc3*tc4*tc5)
}

func treeCountForSlope(world [][]string, stepRight int, stepDown int) int {
	x, y, treeCount := 0, 0, 0
	for y < len(world) {
		if x%(len(world[y])-1) > 0 {
			x = x % (len(world[y]))
		}

		if world[y][x] == "#" {
			treeCount++
		}

		x += stepRight
		y += stepDown
	}

	return treeCount
}

func loadWorld() [][]string {
	data, _ := ioutil.ReadFile("input.txt")
	var input [][]string

	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			input = append(input, strings.Split(line, ""))
		}
	}
	return input
}
