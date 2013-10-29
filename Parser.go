package golumn

import (
	"math"
	"regexp"
	"strconv"
)

type Parser struct {
	lines    []*Line
	padSizes map[int]int
	options  map[string]string
}

func NewParser(input string, options map[string]string) *Parser {
	lines := makeLines(input, options)
	padSizes := makePadSizes(lines)

	return &Parser{
		lines:    lines,
		padSizes: padSizes,
		options:  options,
	}
}

func (p *Parser) Parse(output *string) {
	columnWidth, _ := strconv.Atoi(p.options["columnWidth"])

	for i, line := range p.lines {
		*output += line.Join(p.padSizes, p.options["columnSpacer"], columnWidth)

		if i < len(p.lines)-1 {
			*output += p.options["newLine"]
		}
	}
}

func makeLines(input string, options map[string]string) []*Line {
	lines := regexp.MustCompile(options["newLine"]).Split(input, -1)
	slice := make([]*Line, len(lines))

	for i, line := range lines {
		slice[i] = NewLine(line, options["delim"])
	}

	return slice
}

func makePadSizes(lines []*Line) map[int]int {
	// A map of column-index / pad-size pairs
	padSizes := make(map[int]int)

	for _, line := range lines {
		for columnIndex, chunk := range line.chunks {
			padSizes[columnIndex] = int(math.Max(float64(padSizes[columnIndex]), float64(len(chunk))))
		}
	}

	return padSizes
}
