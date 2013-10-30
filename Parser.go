package golumn

import (
	"math"
	"regexp"
)

type Options struct {
	ColumnWidth    int
	MaxColumnWidth int
	ColumnSpacer   string
	NewLine        string
	Delim          string
}

type Parser struct {
	lines    []*Line
	padSizes map[int]int
	options  *Options
}

func NewParser(input string, options *Options) *Parser {
	lines := makeLines(input, options)
	padSizes := makePadSizes(lines, options)

	return &Parser{
		lines:    lines,
		padSizes: padSizes,
		options:  options,
	}
}

func (p *Parser) Parse(output *string) {
	for i, line := range p.lines {
		*output += line.Join(p.padSizes, p.options.ColumnSpacer, p.options.ColumnWidth)

		if i < len(p.lines)-1 {
			*output += p.options.NewLine
		}
	}
}

func makeLines(input string, options *Options) []*Line {
	lines := regexp.MustCompile(options.NewLine).Split(input, -1)
	slice := make([]*Line, len(lines))

	for i, line := range lines {
		slice[i] = NewLine(line, options.Delim)
	}

	return slice
}

func makePadSizes(lines []*Line, options *Options) map[int]int {
	// A map of column-index / pad-size pairs
	padSizes := make(map[int]int)

	for _, line := range lines {
		for columnIndex, chunk := range line.chunks {
			padSizes[columnIndex] = int(math.Max(float64(padSizes[columnIndex]), float64(len(chunk))))
		}
	}

	return padSizes
}
