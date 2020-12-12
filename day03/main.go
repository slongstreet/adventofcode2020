package main

import (
	"fmt"

	"../common"
)

func main() {
	lines, _ := common.ReadStringsFromFile("input.txt")
	fmt.Printf("read %d lines from file...\n", len(lines))

	slopemap := NewSlopeMap(lines)
	trials := make([]int, 5)

	trials[0] = calculateTreesEncountered(slopemap, 1, 1)
	trials[1] = calculateTreesEncountered(slopemap, 3, 1)
	trials[2] = calculateTreesEncountered(slopemap, 5, 1)
	trials[3] = calculateTreesEncountered(slopemap, 7, 1)
	trials[4] = calculateTreesEncountered(slopemap, 1, 2)

	fmt.Printf("Part One: trees in path = %d\n", trials[1])
	fmt.Printf("Part Two: product of trials = %d\n", trials[0]*trials[1]*trials[2]*trials[3]*trials[4])
}

func calculateTreesEncountered(slopemap *SlopeMap, horizontalRate int, verticalRate int) int {
	numTreesInPath := 0

	var xpos, ypos int
	for ypos = 0; ypos < slopemap.Height; ypos = ypos + verticalRate {
		if slopemap.IsTreeAtPosition(xpos%slopemap.Width, ypos) {
			numTreesInPath++
		}

		xpos = xpos + horizontalRate
	}

	return numTreesInPath
}
