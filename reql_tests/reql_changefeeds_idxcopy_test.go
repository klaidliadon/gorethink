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

// Test duplicate indexes with squashing
func TestChangefeedsIdxcopySuite(t *testing.T) {
	suite.Run(t, new(ChangefeedsIdxcopySuite ))
}

type ChangefeedsIdxcopySuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ChangefeedsIdxcopySuite) SetupTest() {
	suite.T().Log("Setting up ChangefeedsIdxcopySuite")
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

	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *ChangefeedsIdxcopySuite) TearDownSuite() {
	suite.T().Log("Tearing down ChangefeedsIdxcopySuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ChangefeedsIdxcopySuite) TestCases() {
	suite.T().Log("Running ChangefeedsIdxcopySuite: Test duplicate indexes with squashing")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors


	{
		// changefeeds/idxcopy.yaml line #4
		/* partial({'created':1}) */
		var expected_ Expected = partial(map[interface{}]interface{}{"created": 1, })
		/* tbl.index_create('a') */

		suite.T().Log("About to run line #4: tbl.IndexCreate('a')")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreate("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// changefeeds/idxcopy.yaml line #6
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.index_wait('a') */

		suite.T().Log("About to run line #6: tbl.IndexWait('a')")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #6")
	}

	// changefeeds/idxcopy.yaml line #8
	// feed = tbl.order_by(index='a').limit(10).changes(squash=2).limit(9)
	suite.T().Log("Possibly executing: var feed r.Term = tbl.OrderBy(r.OrderByOpts{Index: 'a', }).Limit(10).Changes(r.ChangesOpts{Squash: 2, }).Limit(9)")

	feed := maybeRun(tbl.OrderBy(r.OrderByOpts{Index: "a", }).Limit(10).Changes(r.ChangesOpts{Squash: 2, }).Limit(9), suite.session, r.RunOpts{
		MaxBatchRows: 1,
	});
	_ = feed // Prevent any noused variable errors


	{
		// changefeeds/idxcopy.yaml line #15
		/* partial({'inserted':12, 'errors':0}) */
		var expected_ Expected = partial(map[interface{}]interface{}{"inserted": 12, "errors": 0, })
		/* tbl.insert(r.range(0, 12).map({'id':r.row, 'a':5})) */

		suite.T().Log("About to run line #15: tbl.Insert(r.Range(0, 12).Map(map[interface{}]interface{}{'id': r.Row, 'a': 5, }))")

		runAndAssert(suite.Suite, expected_, tbl.Insert(r.Range(0, 12).Map(map[interface{}]interface{}{"id": r.Row, "a": 5, })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// changefeeds/idxcopy.yaml line #20
		/* partial({'deleted':3, 'errors':0}) */
		var expected_ Expected = partial(map[interface{}]interface{}{"deleted": 3, "errors": 0, })
		/* tbl.get_all(1, 8, 9, index='id').delete() */

		suite.T().Log("About to run line #20: tbl.GetAllByIndex('id', 1, 8, 9).Delete()")

		runAndAssert(suite.Suite, expected_, tbl.GetAllByIndex("id", 1, 8, 9).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// changefeeds/idxcopy.yaml line #26
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* wait(2) */

		suite.T().Log("About to run line #26: wait(2)")

		actual := wait(2)

		assertCompare(suite.T(), expected_, actual)
		suite.T().Log("Finished running line #26")
	}

	{
		// changefeeds/idxcopy.yaml line #28
		/* bag([
{"new_val":{"a":5, "id":0}, "old_val":nil},
{"new_val":{"a":5, "id":2}, "old_val":nil},
{"new_val":{"a":5, "id":3}, "old_val":nil},
{"new_val":{"a":5, "id":4}, "old_val":nil},
{"new_val":{"a":5, "id":5}, "old_val":nil},
{"new_val":{"a":5, "id":6}, "old_val":nil},
{"new_val":{"a":5, "id":7}, "old_val":nil},
{"new_val":{"a":5, "id":10}, "old_val":nil},
{"new_val":{"a":5, "id":11}, "old_val":nil}]) */
		var expected_ Expected = bag([]interface{}{map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 0, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 2, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 3, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 4, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 5, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 6, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 7, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 10, }, "old_val": nil, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 5, "id": 11, }, "old_val": nil, }})
		/* fetch(feed) */

		suite.T().Log("About to run line #28: fetch(feed, 0)")

		fetchAndAssert(suite.Suite, expected_, feed, 0)
		suite.T().Log("Finished running line #28")
	}
}
