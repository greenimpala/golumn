# golumn

Format delimited data into columns. Similar to the unix shell program [`column`](http://linux.about.com/library/cmd/blcmdl1_column.htm).

## Example

Some CSV file:
```
Date,Type,Description,Amount
30/10/2013,DEB,SUPERMARKET 29833,5.15
29/10/2013,DEB,AMAZON MKPLACE,5.63
```
Parse the file:
```go
var result = golumn.parse(string(data), ",", golumn.Options{
	MaxColumnWidth: 5,
	Truncate:       true,
})
print(result)
```
Prints the following:
```
Date 	Type	Descr	Amoun
 30/1	DEB 	SUPER	5.15 
 29/1	DEB 	AMAZO	5.63 
```

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
func ParseF(input string, delim string, options Options) string
```

## Options

For use with `ParseF` - where `options` is a `golumn.Options` struct containing a subset of the following

* `ColumnSpacer` `string` - The characters used to pad columns, default is `\t`.
* `NewLine` `string` - New line character, default is `\n`.
* `ColumnWidth` `int` - Sets colums to be a fixed width.
* `MaxColumnWidth` `int` - Constrains column widths. Overidden if a valid `ColumnWidth` option is given.
* `Truncate` `bool` - Truncates any cells that overflow the `ColumnWidth` or `MaxColumnWidth`.
