# golumn

Pretty print tabular data. Similar to the unix shell program [`column`](http://linux.about.com/library/cmd/blcmdl1_column.htm).

## Usage

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