package main

import (
	"bufio"
	"regexp"

	"strings"

	"fmt"
	"log"
	"os"
	"strconv"
)

type BitNum string

func CreateBitNumFromInt(n int) BitNum {
	stringBitNum := BitNum(strconv.FormatUint(uint64(n), 2)).reverse()
	for i := len(stringBitNum); i < 36; i++ {
		stringBitNum += "0"
	}
	return stringBitNum.reverse()
}

func (b BitNum) MemoryAddresses(addresses []BitNum) []BitNum {
	idx := strings.Index(string(b), "X")

	if idx == -1 {
		addresses = append(addresses, b)
		return addresses
	}

	potential0 := []rune(b)
	potential0[idx] = '0'
	add := BitNum(potential0).MemoryAddresses(addresses)
	addresses = append(addresses, add...)

	potential1 := []rune(b)
	potential1[idx] = '1'
	add = BitNum(potential1).MemoryAddresses(addresses)
	addresses = append(addresses, add...)

	// uniquify
	var ret []BitNum
	for _, addr := range addresses {
		found := false
		for _, r := range ret {
			if r == addr {
				found = true
			}
		}
		if !found {
			ret = append(ret, addr)
		}
	}

	return ret
}

func (b BitNum) ApplyMask(mask BitNum) BitNum {
	out := []rune(b)
	for i := range b {
		if string(mask[i]) == "X" || string(mask[i]) == "1" {
			out[i] = rune(mask[i])
		}
	}
	return BitNum(out)
}

func (b BitNum) ToInt() int {
	i, _ := strconv.ParseInt(string(b), 2, 64)
	return int(i)
}

func (m BitNum) reverse() BitNum {
	runes := []rune(m)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return BitNum(runes)
}

func main() {
	// x := CreateBitNumFromInt(42)
	// var m BitNum = "000000000000000000000000000000X1001X"
	// masked := x.ApplyMask(m)
	// fmt.Println("masked = ", masked)
	var addresses []BitNum
	// fmt.Println(masked.MemoryAddresses(addresses))

	//fmt.Println(m.MemoryAddresses())
	var currentMask BitNum
	memory := make(map[int]int)

	r := regexp.MustCompile(`\d+`)
	for _, line := range readInput() {
		if strings.HasPrefix(line, "mask") {
			currentMask = BitNum(strings.Split(line, " = ")[1])
		} else {
			matches := r.FindAllString(line, -1)
			addr, _ := strconv.Atoi(matches[0])
			val, _ := strconv.Atoi(matches[1])

			// fmt.Println("applying to = ", CreateBitNumFromInt(addr))
			// fmt.Println("currentMaks = ", currentMask)
			maskedAddr := CreateBitNumFromInt(addr).ApplyMask(currentMask)
			// fmt.Println("result      = ", maskedAddr)
			// fmt.Println("addr = ", addr, ", val = ", val, ", maskedAddr = ", maskedAddr)

			for _, a := range maskedAddr.MemoryAddresses(addresses) {
				// fmt.Println("a.ToInt = ", a.ToInt(), ", val = ", val)
				memory[a.ToInt()] = val
			}
		}
	}

	sum := 0
	for _, mem := range memory {
		sum += mem
	}
	fmt.Println("Sum = ", sum)
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output
}
