package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// https://www.geeksforgeeks.org/determine-the-count-of-leaf-nodes-in-an-n-ary-tree/

type Node struct {
	Val    int
	Leaves []*Node
}

var endpoints int

//  oneRuns groups the number of 1 diffs together, so [1,1,1] will become [3].
func totals(arr []int) []int {
	var diffs []int
	count := 0
	for i := range arr {
		if i < len(arr)-1 && arr[i+1]-arr[i] == 1 {
			count++
		} else {
			if count > 0 {
				diffs = append(diffs, count)
			}
			count = 0
		}
	}

	var totals []int
	for i := range diffs {
		n := diffs[i]
		totals = append(totals, 1+(n*(n-1)/2))
	}
	return totals
}

func main() {
	adapters := readInput()
	adapters = append(adapters, 0)

	sort.Ints(adapters)

	t := totals(adapters)
	product := 1
	for _, n := range t {
		product *= n
	}
	fmt.Println("product = ", product)
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
