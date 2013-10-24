package golumn

import (
	"math"
	"regexp"
)

const (
	columnSpacer = "\t"
	newLine      = "\n"
	space        = " "
)

func Parse(input string, delim string) (output string) {
	if delim == "" {
		return input
	}

	lines := lines(input)
	padSizes := determinePadSizes(lines, delim)

	for i, line := range lines {
		line = padChunks(line, delim, padSizes)
		line = replaceDelimiter(line, delim)

		output += line

		if i < len(lines)-1 {
			output += newLine
		}
	}
	return
}

func determinePadSizes(lines []string, delim string) map[int]int {
	// A map of column-index / pad-size pairs
	padSizes := make(map[int]int)

	for _, line := range lines {
		chunks := chunks(line, delim)

		for columnIndex, chunk := range chunks {
			padSizes[columnIndex] = int(math.Max(float64(padSizes[columnIndex]), float64(len(chunk))))
		}
	}

	return padSizes
}

func padChunks(line string, delim string, padSize map[int]int) (output string) {
	chunks := chunks(line, delim)

	for columnIndex, chunk := range chunks {
		// While chunk is less than
		// max width for this column
		for len(chunk) < padSize[columnIndex] {
			chunk += space
		}

		output += chunk

		// Append delimiter if not final chunk
		if columnIndex < len(chunks)-1 {
			output += delim
		}
	}

	return
}

func replaceDelimiter(line string, delim string) string {
	return regexp.MustCompile(delim).ReplaceAllString(line, columnSpacer)
}

func lines(input string) []string {
	return regexp.MustCompile(newLine).Split(input, -1)
}

func chunks(line string, delim string) []string {
	return regexp.MustCompile(delim).Split(line, -1)
}
