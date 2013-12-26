package golumn_test

import (
	"github.com/st3redstripe/golumn"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestReturnsOriginalWithBlankDelimiter(c *C) {
	c.Check(golumn.Parse("foo", ""), Equals, "foo")
}

func (s *MySuite) TestCanParseSingleLine(c *C) {
	c.Check(golumn.Parse("a,b,c", ","), Equals, "a\tb\tc")
}

func (s *MySuite) TestLargeDelimiter(c *C) {
	c.Check(golumn.Parse("agolangbgolangc", "golang"), Equals, "a\tb\tc")
}

func (s *MySuite) TestSameColumnSizeMultiLine(c *C) {
	input := "a,b,c\na,b,c"
	c.Check(golumn.Parse(input, ","), Equals, "a\tb\tc\na\tb\tc")
}

func (s *MySuite) TestVariableColumnSizeMultiLine(c *C) {
	input := "aaa,b,cc\na,bb,c"

	c.Check(golumn.Parse(input, ","), Equals, "aaa\tb \tcc\na  \tbb\tc ")
}

func (s *MySuite) TestParseFCanTakeCustomColumnSpacer(c *C) {
	input := "aaa,b,cc\na,bb,c"
	options := golumn.Options{
		ColumnSpacer: "**",
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aaa**b **cc\na  **bb**c ")
}

func (s *MySuite) TestParseFCanTakeCustomNewLineCharacter(c *C) {
	input := "aaa,b,cc\ra,bb,c"
	options := golumn.Options{
		NewLine: "\r",
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aaa\tb \tcc\ra  \tbb\tc ")
}

func (s *MySuite) TestParseFCanFormatForCustomColumnWidthSingleLine(c *C) {
	input := "aaa,bbb,ccc"
	options := golumn.Options{
		ColumnWidth: 2,
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aa\tbb\tcc\na \tb \tc ")
}

func (s *MySuite) TestParseFCanFormatForCustomColumnWidthMulti(c *C) {
	input := "aaa,bbb,ccc\nzz,xxxx,yy"
	options := golumn.Options{
		ColumnWidth: 2,
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aa\tbb\tcc\na \tb \tc \nzz\txx\tyy\n  \txx\t  ")
}

func (s *MySuite) TestParseFCanFormatForMaxWidth(c *C) {
	input := "aa,bbbb,cccc"
	options := golumn.Options{
		MaxColumnWidth: 2,
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aa\tbb\tcc\n  \tbb\tcc")
}

func (s *MySuite) TestParseFCanTruncate(c *C) {
	input := "aa,bbbb,cccc"
	options := golumn.Options{
		ColumnWidth: 2,
		Truncate:    true,
	}

	c.Check(golumn.ParseF(input, ",", options), Equals, "aa\tbb\tcc")
}
