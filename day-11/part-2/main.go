package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Layout struct {
	Seats [][]string
}

func (l *Layout) OccupySeats() {
	newSeats := l.duplicateSeats()
	for y := 0; y < len(l.Seats); y++ {
		for x := 0; x < len(l.Seats); x++ {
			if newSeats[y][x] != "L" {
				continue
			}

			if l.adjacentCount(x, y) == 0 {
				newSeats[y][x] = "#"
			}
		}
	}
	l.Seats = newSeats
}

func (l *Layout) EmptySeats() {
	newSeats := l.duplicateSeats()
	for y := 0; y < len(l.Seats); y++ {
		for x := 0; x < len(l.Seats); x++ {
			if l.Seats[y][x] != "#" {
				continue
			}
			if l.adjacentCount(x, y) >= 5 {
				newSeats[y][x] = "L"
			}
		}
	}
	l.Seats = newSeats
}

func (l *Layout) OccupiedTotal() int {
	total := 0
	for _, row := range l.Seats {
		for _, char := range row {
			if char == "#" {
				total++
			}
		}
	}
	return total
}

func (l *Layout) duplicateSeats() [][]string {
	dup := make([][]string, len(l.Seats))
	for i := range l.Seats {
		dup[i] = make([]string, len(l.Seats))
		copy(dup[i], l.Seats[i])
	}
	return dup
}

func (l *Layout) adjacentCount(x int, y int) int {
	count := 0
	if l.occupiedSeatInDirection(x, y, -1, -1) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, 0, -1) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, 1, -1) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, -1, 0) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, 1, 0) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, -1, 1) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, 0, 1) {
		count++
	}
	if l.occupiedSeatInDirection(x, y, 1, 1) {
		count++
	}

	return count
}

func (l *Layout) occupiedSeatInDirection(x int, y int, iterX int, iterY int) bool {
	for {
		x += iterX
		y += iterY

		if y < 0 || x < 0 || y >= len(l.Seats) || x >= len(l.Seats[y]) {
			return false
		}
		if l.Seats[y][x] == "L" {
			return false
		}
		if l.Seats[y][x] == "#" {
			return true
		}
	}
}

func (l *Layout) printSeats() {
	for _, row := range l.Seats {
		fmt.Println(row)
	}
}

func main() {
	layout := readInput()
	for x := 0; x < 100; x++ {
		layout.OccupySeats()
		layout.EmptySeats()
		fmt.Println("Total: ", layout.OccupiedTotal())
	}
}

func readInput() *Layout {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	layout := &Layout{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		layout.Seats = append(layout.Seats, strings.Split(scanner.Text(), ""))
	}

	return layout
}
