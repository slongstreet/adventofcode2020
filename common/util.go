package common

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ReadStringsFromFile : Read an array of strings from lines in a file
func ReadStringsFromFile(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return lines, nil
}

// ReadIntegersFromFile : Read an array of integers from lines in a file
func ReadIntegersFromFile(inputFile string) ([]int, error) {
	lines, err := ReadStringsFromFile(inputFile)
	if err != nil {
		return nil, err
	}

	ints := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Infof("unable to parse string as int: %s (%s)", line, err)
			continue
		}

		ints[i] = num
	}

	return ints, nil
}
