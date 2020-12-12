package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

type Ship struct {
	Orientation string
	X           int
	Y           int
}

func (s *Ship) Move(dir string, arg int) {
	if dir == "F" {
		dir = s.Orientation
	}
	switch {
	case dir == "N":
		s.Y -= arg
	case dir == "S":
		s.Y += arg
	case dir == "E":
		s.X -= arg
	case dir == "W":
		s.X += arg
	case dir == "R":
		s.changeDirection(dir, arg)
	case dir == "L":
		s.changeDirection(dir, arg)
	}
}

func (s *Ship) changeDirection(dir string, arg int) {
	fmt.Println("-------")
	fmt.Println("change direction", dir, arg)
	dirs := []string{"N", "E", "S", "W"}
	idx := 0
	for i, d := range dirs {
		if d == s.Orientation {
			idx = i
		}
	}
	ticks := arg / 90

	fmt.Println("idx = ", idx)
	fmt.Println("ticks = ", ticks)

	if dir == "R" {
		idx += ticks
		if idx >= 4 {
			idx %= 4
		}
		s.Orientation = dirs[idx]
	} else {
		for n := 0; n < ticks; n++ {
			idx--
			if idx == -1 {
				idx = 3
			}
		}
		s.Orientation = dirs[idx]
	}
}

func (s *Ship) ManhattanDistance() int {
	return abs(s.X) + abs(s.Y)
}

func (s *Ship) Print() {
	fmt.Printf("(%d, %d), orientation: %s\n", s.X, s.Y, s.Orientation)
}

func main() {
	dirs := readInput()
	ship := &Ship{Orientation: "E", X: 0, Y: 0}
	for _, dir := range dirs {
		direction := string(dir[0])
		val, _ := strconv.Atoi(string(dir[1:]))
		ship.Move(direction, val)
		ship.Print()
	}
	fmt.Println("Manhattan: ", ship.ManhattanDistance())
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dirs []string
	for scanner.Scan() {
		dirs = append(dirs, scanner.Text())
	}

	return dirs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
