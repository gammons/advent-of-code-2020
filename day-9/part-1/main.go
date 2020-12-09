package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const Preamble = 25

func main() {
	allnums := readInput()
	for i := Preamble + 1; i < len(allnums); i++ {
		consider := make([]int, Preamble)
		copy(consider, allnums[i-Preamble:i])
		sort.Ints(consider)
		if !isValid(allnums[i], consider) {
			fmt.Printf("%d is invalid\n", allnums[i])
			break
		}
	}
}

func isValid(num int, consider []int) bool {
	for x := range consider {
		for y := x; y < len(consider); y++ {
			if consider[x]+consider[y] == num {
				return true
			}
		}
	}
	return false
}

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ints []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, num)
	}

	return ints
}
