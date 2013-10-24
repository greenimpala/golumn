# golumn

Pretty print tabular data. Similar to the unix shell program [`column`](http://linux.about.com/library/cmd/blcmdl1_column.htm).

## Usage

Call `Parse` passing in an input string and a delimiter.

```go
func Parse(input string, delim string) string
```