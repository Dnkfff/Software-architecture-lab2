package lab2

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestExpressionToPostfixSum(c *C) {
	got, _ := ExpressionToPostfix("100+100")

	c.Assert(got, Equals, "100 100 +")
}

func (s *TestSuite) TestExpressionToPostfixMult(c *C) {
	got, _ := ExpressionToPostfix("100*100")

	c.Assert(got, Equals, "100 100 *")
}

func (s *TestSuite) TestExpressionToPostfixEmptyString(c *C) {
	got, _ := ExpressionToPostfix("")

	c.Assert(got, Equals, "")
}

func (s *TestSuite) TestExpressionToPostfixComplicated(c *C) {
	got, _ := ExpressionToPostfix("(123-123)^6*(56-22)")

	c.Assert(got, Equals, "123 123 - 6 ^ 56 22 - *")
	// Output:
	// 123 123 - 6 ^ 56 22 - *
}

func (s *TestSuite) TestExpressionToPostfixComplicated2(c *C) {
	got, _ := ExpressionToPostfix("5+(4-2)*3")

	c.Assert(got, Equals, "5 4 2 - 3 * +")

	// Output:
	// 5 4 2 - 3 * +
}
