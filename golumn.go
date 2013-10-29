package golumn

var defaultOptions = map[string]string{
	"columnSpacer": "\t",
	"newLine":      "\n",
}

func Parse(input string, delim string) string {
	return parse(input, delim, defaultOptions)
}

func ParseF(input string, delim string, options map[string]string) string {
	// Mix-in defaults for non-existant keys
	for key, value := range defaultOptions {
		if options[key] == "" {
			options[key] = value
		}
	}

	return parse(input, delim, options)
}

func parse(input string, delim string, options map[string]string) (output string) {
	if delim == "" {
		return input
	}
	options["delim"] = delim

	parser := NewParser(input, options)
	parser.Parse(&output)

	return
}
