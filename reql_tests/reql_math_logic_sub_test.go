// Autogenerated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go.tpl
package reql_tests

import (
"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

// Tests for basic usage of the subtraction operation
func TestMathLogicSubSuite(t *testing.T) {
	suite.Run(t, new(MathLogicSubSuite ))
}

type MathLogicSubSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicSubSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicSubSuite")
	// Use imports to prevent errors
	time.Now()

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicSubSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicSubSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicSubSuite) TestCases() {
	suite.T().Log("Running MathLogicSubSuite: Tests for basic usage of the subtraction operation")



	{
		// math_logic/sub.yaml line #6
		/* 0 */
		var expected_ int = 0
		/* r.expr(1) - 1 */

		suite.T().Log("About to run line #6: r.Expr(1).Sub(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Sub(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// math_logic/sub.yaml line #7
		/* 0 */
		var expected_ int = 0
		/* 1 - r.expr(1) */

		suite.T().Log("About to run line #7: r.Sub(1, r.Expr(1))")

		runAndAssert(suite.Suite, expected_, r.Sub(1, r.Expr(1)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// math_logic/sub.yaml line #8
		/* 0 */
		var expected_ int = 0
		/* r.expr(1).sub(1) */

		suite.T().Log("About to run line #8: r.Expr(1).Sub(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Sub(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/sub.yaml line #17
		/* -2 */
		var expected_ int = -2
		/* r.expr(-1) - 1 */

		suite.T().Log("About to run line #17: r.Expr(-1).Sub(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(-1).Sub(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// math_logic/sub.yaml line #22
		/* -6.75 */
		var expected_ float64 = -6.75
		/* r.expr(1.75) - 8.5 */

		suite.T().Log("About to run line #22: r.Expr(1.75).Sub(8.5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.75).Sub(8.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// math_logic/sub.yaml line #26
		/* -13 */
		var expected_ int = -13
		/* r.expr(1).sub(2,3,4,5) */

		suite.T().Log("About to run line #26: r.Expr(1).Sub(2, 3, 4, 5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Sub(2, 3, 4, 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// math_logic/sub.yaml line #30
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('a').sub(0.8) */

		suite.T().Log("About to run line #30: r.Expr('a').Sub(0.8)")

		runAndAssert(suite.Suite, expected_, r.Expr("a").Sub(0.8), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// math_logic/sub.yaml line #33
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr(1).sub('a') */

		suite.T().Log("About to run line #33: r.Expr(1).Sub('a')")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Sub("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #33")
	}

	{
		// math_logic/sub.yaml line #36
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('b').sub('a') */

		suite.T().Log("About to run line #36: r.Expr('b').Sub('a')")

		runAndAssert(suite.Suite, expected_, r.Expr("b").Sub("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #36")
	}
}
