package main

const open rune = '.'
const tree rune = '#'

// SlopeMap represents the sloped hillside with occasional trees
type SlopeMap struct {
	Width     int
	Height    int
	locations [][]rune
}

// NewSlopeMap creates a new SlopeMap initialized with the supplied input
func NewSlopeMap(input []string) *SlopeMap {
	slopemap := &SlopeMap{
		Height:    len(input),
		locations: make([][]rune, 0),
	}

	for _, line := range input {
		slopemap.locations = append(slopemap.locations, []rune(line))
	}

	slopemap.Width = len(slopemap.locations[0])

	return slopemap
}

// IsTreeAtPosition determines whether there is a tree at the given position
func (slopemap *SlopeMap) IsTreeAtPosition(x int, y int) bool {
	return slopemap.locations[y][x] == tree
}
