# golumn

Format delimited data into columns. Similar to the unix shell program [`column`](http://linux.about.com/library/cmd/blcmdl1_column.htm).

## Usage

Download and compile the source with `go get github.com/st3redstripe/golumn`. Then import as usual.

```go
import (
	"github.com/st3redstripe/golumn"
)
```

Call `Parse` passing in an input string and a delimiter.

```go
func Parse(input string, delim string) string
```

Use `ParseF` to override default options.

```go
func ParseF(input string, delim string, options map[string]string) string
```

Where options is a combination of the following:

* `columnSpacer` - The characters used to pad columns, default is `\t`.
* `newLine` - New line character, default is `\n`.
* `columnWidth` - Sets colums to be a fixed width.
