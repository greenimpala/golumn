package golumn

import (
	"math"
	"regexp"
	"strings"
)

type Line struct {
	chunks []string
}

func NewLine(line string, delim string) *Line {
	return &Line{
		chunks: regexp.MustCompile(delim).Split(line, -1),
	}
}

func (l *Line) Join(padSizes map[int]int, delim string, columnWidth int) string {
	if columnWidth > 0 {
		return l.joinWithColumnConstraints(delim, columnWidth)
	} else {
		return l.joinWithPadding(delim, padSizes)
	}
}

func (l *Line) joinWithPadding(delim string, padSizes map[int]int) string {
	chunks := make([]string, len(l.chunks))

	for i, chunk := range l.chunks {
		for len(chunk) < padSizes[i] {
			chunk += " "
		}
		chunks[i] = chunk
	}

	return strings.Join(chunks, delim)
}

func (l *Line) joinWithColumnConstraints(delim string, columnWidth int) (output string) {
	var lines float64 // Number of lines we are going to fill

	// Calculate the max lines
	for _, chunk := range l.chunks {
		lines = math.Max(lines, math.Ceil(float64(len(chunk))/float64(columnWidth)))
	}

	chunksBuffer := l.chunks

	// Loop through each line and grab as many chars from
	// the chunks buffer as will fit in the column
	for i := 0; i < int(lines); i++ {
		for columnIndex, chunk := range chunksBuffer {
			// Pad chunk
			for len(chunk) < columnWidth {
				chunk += " "
			}

			output += chunk[:columnWidth]
			chunksBuffer[columnIndex] = chunk[columnWidth:]

			if columnIndex < len(chunksBuffer)-1 {
				output += delim
			}
		}
		if i < int(lines)-1 {
			output += "\n"
		}
	}
	return
}
