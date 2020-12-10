package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := readInput()
	sort.Ints(adapters)
	oneDiff, threeDiff := 1, 1

	for i := 0; i < len(adapters)-1; i++ {
		if adapters[i+1]-adapters[i] == 1 {
			oneDiff++
		}
		if adapters[i+1]-adapters[i] == 3 {
			threeDiff++
		}
	}
	fmt.Println(oneDiff * threeDiff)
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
