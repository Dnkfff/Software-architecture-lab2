package lab2

import (
	"bytes"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestInputFile(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("5+5"),
		Output: &buff,
	}
	var expected = "5 5 +"

	handler.Compute()
	c.Assert(buff.String(), Equals, expected)
}

func (s *TestSuite) TestInputErr(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("a+c"),
		Output: &buff,
	}
	expected := "the string is invalid"

	err := handler.Compute()
	c.Assert(fmt.Sprint(err), Equals, expected)
}
