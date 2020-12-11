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
			if l.adjacentCount(x, y) >= 4 {
				newSeats[y][x] = "L"
			}
		}
	}
	l.Seats = newSeats
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
	// row above
	if y > 0 {
		if x > 0 && l.Seats[y-1][x-1] == "#" {
			count++
		}

		if l.Seats[y-1][x] == "#" {
			count++
		}

		if x < len(l.Seats[y])-1 && l.Seats[y-1][x+1] == "#" {
			count++
		}
	}

	// current row
	if x > 0 && l.Seats[y][x-1] == "#" {
		count++
	}

	if x < len(l.Seats[y])-1 && l.Seats[y][x+1] == "#" {
		count++
	}

	// row below
	if y < len(l.Seats)-1 {
		if x > 0 && l.Seats[y+1][x-1] == "#" {
			count++
		}

		if l.Seats[y+1][x] == "#" {
			count++
		}

		if x < len(l.Seats[y])-1 && l.Seats[y+1][x+1] == "#" {
			count++
		}
	}
	return count
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
