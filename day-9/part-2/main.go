package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// 22406676 is my sum
func main() {
	allnums := readInput()
	for i := 0; i < len(allnums); i++ {
		sum := allnums[i]
		for j := i + 1; j < len(allnums); j++ {
			sum += allnums[j]
			if sum == 22406676 {
				arr := make([]int, (j-i)+1)
				copy(arr, allnums[i:j+1])
				sort.Ints(arr)
				fmt.Printf("Sum found: %d\n", arr[0]+arr[len(arr)-1])
			}
			if sum > 22406676 {
				break
			}
		}
	}
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
