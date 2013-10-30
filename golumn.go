package golumn

var defaultOptions = Options{
	ColumnSpacer: "\t",
	NewLine:      "\n",
	Delim:        ",",
}

func Parse(input string, delim string) string {
	options := defaultOptions
	options.Delim = delim
	return parse(input, options)
}

func ParseF(input string, delim string, options Options) string {
	if options.ColumnSpacer == "" {
		options.ColumnSpacer = defaultOptions.ColumnSpacer
	}
	if options.NewLine == "" {
		options.NewLine = defaultOptions.NewLine
	}
	options.Delim = delim

	return parse(input, options)
}

func parse(input string, options Options) (output string) {
	if options.Delim == "" {
		return input
	}

	parser := NewParser(input, &options)
	parser.Parse(&output)

	return
}
