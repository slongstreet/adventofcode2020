package main

import (
	"fmt"
	"strings"

	"../common"
)

func main() {
	lines, _ := common.ReadStringsFromFile("input.txt")
	fmt.Printf("read %d lines from file...\n", len(lines))

	validPasswordsPartOne := 0
	validPasswordsPartTwo := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		policy := NewPasswordPolicy(parts[0])
		if policy.IsValidPasswordPartOne(strings.Trim(parts[1], " ")) {
			validPasswordsPartOne++
		}
		if policy.IsValidPasswordPartTwo(strings.Trim(parts[1], " ")) {
			validPasswordsPartTwo++
		}
	}

	fmt.Printf("Part One: valid password count = %d\n", validPasswordsPartOne)
	fmt.Printf("Part Two: valid password count = %d\n", validPasswordsPartTwo)
}
