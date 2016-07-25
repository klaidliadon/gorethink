// Autogenerated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMathLogicMathSuite(t *testing.T) {
    suite.Run(t, new(MathLogicMathSuite ))
}

type MathLogicMathSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicMathSuite) SetupTest() {
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

func (suite *MathLogicMathSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *MathLogicMathSuite) TestCases() {


    {
        // math_logic/math.yaml line #4
        /* 1 */
        var expected_ int = 1
        /* (((4 + 2 * (r.expr(26) % 18)) / 5) - 3) */

    	suite.T().Log("About to run line #4: r.Add(4, r.Mul(2, r.Expr(26).Mod(18))).Div(5).Sub(3)")

        cursor, err := r.Add(4, r.Mul(2, r.Expr(26).Mod(18))).Div(5).Sub(3).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #4")
    }
}