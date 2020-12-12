package main

import (
	"fmt"

	"../common"
)

func main() {
	ints, _ := common.ReadIntegersFromFile("input.txt")
	fmt.Printf("read %d integers from file...\n", len(ints))

	fmt.Printf("Part One: Product = %d\n", partOne(ints))
	fmt.Printf("Part Two: Product = %d\n", partTwo(ints))
}

func partOne(ints []int) int {
	for i, a := range ints {
		for j, b := range ints {
			if i == j {
				continue
			}
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func partTwo(ints []int) int {
	for i, a := range ints {
		for j, b := range ints {
			for k, c := range ints {
				if i == j || j == k {
					continue
				}
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}
