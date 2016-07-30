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

// Test `include_states`
func TestChangefeedsIncludeStatesSuite(t *testing.T) {
	suite.Run(t, new(ChangefeedsIncludeStatesSuite ))
}

type ChangefeedsIncludeStatesSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ChangefeedsIncludeStatesSuite) SetupTest() {
	suite.T().Log("Setting up ChangefeedsIncludeStatesSuite")
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

func (suite *ChangefeedsIncludeStatesSuite) TearDownSuite() {
	suite.T().Log("Tearing down ChangefeedsIncludeStatesSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ChangefeedsIncludeStatesSuite) TestCases() {
	suite.T().Log("Running ChangefeedsIncludeStatesSuite: Test `include_states`")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors


	{
		// changefeeds/include_states.yaml line #4
		/* [{'state':'ready'}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "ready", }}
		/* tbl.changes(squash=true, include_states=true).limit(1) */

		suite.T().Log("About to run line #4: tbl.Changes(r.ChangesOpts{Squash: true, IncludeStates: true, }).Limit(1)")

		runAndAssert(suite.Suite, expected_, tbl.Changes(r.ChangesOpts{Squash: true, IncludeStates: true, }).Limit(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// changefeeds/include_states.yaml line #9
		/* [{'state':'initializing'}, {'new_val':null}, {'state':'ready'}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"new_val": nil, }, map[interface{}]interface{}{"state": "ready", }}
		/* tbl.get(0).changes(squash=true, include_states=true, include_initial=true).limit(3) */

		suite.T().Log("About to run line #9: tbl.Get(0).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(3)")

		runAndAssert(suite.Suite, expected_, tbl.Get(0).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// changefeeds/include_states.yaml line #14
		/* [{'state':'initializing'}, {'state':'ready'}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"state": "ready", }}
		/* tbl.order_by(index='id').limit(10).changes(squash=true, include_states=true, include_initial=true).limit(2) */

		suite.T().Log("About to run line #14: tbl.OrderBy(r.OrderByOpts{Index: 'id', }).Limit(10).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(2)")

		runAndAssert(suite.Suite, expected_, tbl.OrderBy(r.OrderByOpts{Index: "id", }).Limit(10).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// changefeeds/include_states.yaml line #19
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.insert({'id':1}) */

		suite.T().Log("About to run line #19: tbl.Insert(map[interface{}]interface{}{'id': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 1, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// changefeeds/include_states.yaml line #21
		/* [{'state':'initializing'}, {'new_val':{'id':1}}, {'state':'ready'}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"id": 1, }, }, map[interface{}]interface{}{"state": "ready", }}
		/* tbl.order_by(index='id').limit(10).changes(squash=true, include_states=true, include_initial=true).limit(3) */

		suite.T().Log("About to run line #21: tbl.OrderBy(r.OrderByOpts{Index: 'id', }).Limit(10).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(3)")

		runAndAssert(suite.Suite, expected_, tbl.OrderBy(r.OrderByOpts{Index: "id", }).Limit(10).Changes(r.ChangesOpts{Squash: true, IncludeStates: true, IncludeInitial: true, }).Limit(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #21")
	}

	// changefeeds/include_states.yaml line #26
	// tblchanges = tbl.changes(squash=true, include_states=true)
	suite.T().Log("Possibly executing: var tblchanges r.Term = tbl.Changes(r.ChangesOpts{Squash: true, IncludeStates: true, })")

	tblchanges := maybeRun(tbl.Changes(r.ChangesOpts{Squash: true, IncludeStates: true, }), suite.session, r.RunOpts{
	});
	_ = tblchanges // Prevent any noused variable errors


	{
		// changefeeds/include_states.yaml line #30
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.insert({'id':2}) */

		suite.T().Log("About to run line #30: tbl.Insert(map[interface{}]interface{}{'id': 2, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 2, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// changefeeds/include_states.yaml line #32
		/* [{'state':'ready'},{'new_val':{'id':2},'old_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "ready", }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"id": 2, }, "old_val": nil, }}
		/* fetch(tblchanges, 2) */

		suite.T().Log("About to run line #32: fetch(tblchanges, 2)")

		fetchAndAssert(suite.Suite, expected_, tblchanges, 2)
		suite.T().Log("Finished running line #32")
	}

	// changefeeds/include_states.yaml line #35
	// getchanges = tbl.get(2).changes(include_states=true, include_initial=true)
	suite.T().Log("Possibly executing: var getchanges r.Term = tbl.Get(2).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, })")

	getchanges := maybeRun(tbl.Get(2).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, }), suite.session, r.RunOpts{
	});
	_ = getchanges // Prevent any noused variable errors


	{
		// changefeeds/include_states.yaml line #39
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.get(2).update({'a':1}) */

		suite.T().Log("About to run line #39: tbl.Get(2).Update(map[interface{}]interface{}{'a': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(2).Update(map[interface{}]interface{}{"a": 1, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// changefeeds/include_states.yaml line #41
		/* [{'state':'initializing'}, {'new_val':{'id':2}}, {'state':'ready'}, {'old_val':{'id':2},'new_val':{'id':2,'a':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"id": 2, }, }, map[interface{}]interface{}{"state": "ready", }, map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 2, }, "new_val": map[interface{}]interface{}{"id": 2, "a": 1, }, }}
		/* fetch(getchanges, 4) */

		suite.T().Log("About to run line #41: fetch(getchanges, 4)")

		fetchAndAssert(suite.Suite, expected_, getchanges, 4)
		suite.T().Log("Finished running line #41")
	}

	// changefeeds/include_states.yaml line #44
	// limitchanges = tbl.order_by(index='id').limit(10).changes(include_states=true, include_initial=true)
	suite.T().Log("Possibly executing: var limitchanges r.Term = tbl.OrderBy(r.OrderByOpts{Index: 'id', }).Limit(10).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, })")

	limitchanges := maybeRun(tbl.OrderBy(r.OrderByOpts{Index: "id", }).Limit(10).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, }), suite.session, r.RunOpts{
	});
	_ = limitchanges // Prevent any noused variable errors


	// changefeeds/include_states.yaml line #48
	// limitchangesdesc = tbl.order_by(index=r.desc('id')).limit(10).changes(include_states=true, include_initial=true)
	suite.T().Log("Possibly executing: var limitchangesdesc r.Term = tbl.OrderBy(r.OrderByOpts{Index: r.Desc('id'), }).Limit(10).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, })")

	limitchangesdesc := maybeRun(tbl.OrderBy(r.OrderByOpts{Index: r.Desc("id"), }).Limit(10).Changes(r.ChangesOpts{IncludeStates: true, IncludeInitial: true, }), suite.session, r.RunOpts{
	});
	_ = limitchangesdesc // Prevent any noused variable errors


	{
		// changefeeds/include_states.yaml line #52
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.insert({'id':3}) */

		suite.T().Log("About to run line #52: tbl.Insert(map[interface{}]interface{}{'id': 3, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 3, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// changefeeds/include_states.yaml line #54
		/* [{'state':'initializing'}, {'new_val':{'id':1}}, {'new_val':{'a':1, 'id':2}}, {'state':'ready'}, {'old_val':null, 'new_val':{'id':3}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"id": 1, }, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 1, "id": 2, }, }, map[interface{}]interface{}{"state": "ready", }, map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 3, }, }}
		/* fetch(limitchanges, 5) */

		suite.T().Log("About to run line #54: fetch(limitchanges, 5)")

		fetchAndAssert(suite.Suite, expected_, limitchanges, 5)
		suite.T().Log("Finished running line #54")
	}

	{
		// changefeeds/include_states.yaml line #57
		/* [{'state':'initializing'}, {'new_val':{'a':1, 'id':2}}, {'new_val':{'id':1}}, {'state':'ready'}, {'old_val':null, 'new_val':{'id':3}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"state": "initializing", }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"a": 1, "id": 2, }, }, map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"id": 1, }, }, map[interface{}]interface{}{"state": "ready", }, map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 3, }, }}
		/* fetch(limitchangesdesc, 5) */

		suite.T().Log("About to run line #57: fetch(limitchangesdesc, 5)")

		fetchAndAssert(suite.Suite, expected_, limitchangesdesc, 5)
		suite.T().Log("Finished running line #57")
	}
}
