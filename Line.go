package golumn

import (
	"math"
	"regexp"
)

type Line struct {
	chunks []string
}

func NewLine(line string, delim string) *Line {
	return &Line{
		chunks: regexp.MustCompile(delim).Split(line, -1),
	}
}

func (l *Line) Join(padSizes map[int]int, delim string, truncate bool) (output string) {
	var lines float64 // Number of lines we are going to fill

	if truncate {
		lines = 1
	} else {
		// Calculate the max lines from the pad sizes
		for i, chunk := range l.chunks {
			lines = math.Max(lines, math.Ceil(float64(len(chunk))/float64(padSizes[i])))
		}
	}

	chunksBuffer := l.chunks

	// Loop through each line and grab as many chars from
	// the chunks buffer as will fit in the column
	for i := 0; i < int(lines); i++ {
		for columnIndex, chunk := range chunksBuffer {
			// Pad chunk
			for len(chunk) < padSizes[columnIndex] {
				chunk += " "
			}

			if truncate {
				chunk = chunk[:padSizes[columnIndex]]
			}

			output += chunk[:padSizes[columnIndex]]
			chunksBuffer[columnIndex] = chunk[padSizes[columnIndex]:]

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
