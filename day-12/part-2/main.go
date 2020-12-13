package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

type Ship struct {
	X         int
	Y         int
	WaypointX int
	WaypointY int
}

func (s *Ship) Move(dir string, arg int) {
	switch {
	case dir == "F":
		s.X += s.WaypointX * arg
		s.Y += s.WaypointY * arg
	case dir == "N":
		s.WaypointY += arg
	case dir == "S":
		s.WaypointY += arg * -1
	case dir == "E":
		s.WaypointX += arg
	case dir == "W":
		s.WaypointX += arg * -1
	case dir == "R":
		s.changeWaypointDirection(dir, arg)
	case dir == "L":
		s.changeWaypointDirection(dir, arg)
	}
}

func (s *Ship) changeWaypointDirection(dir string, arg int) {
	for n := 0; n < (arg / 90); n++ {
		if dir == "R" {
			s.rotateWaypointClockwise()
		} else {
			s.rotateWaypointCounterclockwise()
		}
	}
}

func (s *Ship) rotateWaypointClockwise() {
	origX := s.WaypointX
	s.WaypointX = s.WaypointY
	s.WaypointY = -1 * origX
}

func (s *Ship) rotateWaypointCounterclockwise() {
	origY := s.WaypointY
	s.WaypointY = s.WaypointX
	s.WaypointX = -1 * origY
}

func (s *Ship) ManhattanDistance() int {
	return abs(s.X) + abs(s.Y)
}

func main() {
	ship := &Ship{X: 0, Y: 0, WaypointX: 10, WaypointY: 1}

	for _, dir := range readInput() {
		direction := string(dir[0])
		val, _ := strconv.Atoi(string(dir[1:]))
		ship.Move(direction, val)
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
