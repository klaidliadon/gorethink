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

// Tests for ordered_union
func TestTransformUnorderedMapSuite(t *testing.T) {
	suite.Run(t, new(TransformUnorderedMapSuite ))
}

type TransformUnorderedMapSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TransformUnorderedMapSuite) SetupTest() {
	suite.T().Log("Setting up TransformUnorderedMapSuite")
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

	r.DB("test").TableDrop("even").Exec(suite.session)
	err = r.DB("test").TableCreate("even").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("even").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("odd").Exec(suite.session)
	err = r.DB("test").TableCreate("odd").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("odd").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("odd2").Exec(suite.session)
	err = r.DB("test").TableCreate("odd2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("odd2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *TransformUnorderedMapSuite) TearDownSuite() {
	suite.T().Log("Tearing down TransformUnorderedMapSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("even").Exec(suite.session)
		 r.DB("test").TableDrop("odd").Exec(suite.session)
		 r.DB("test").TableDrop("odd2").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TransformUnorderedMapSuite) TestCases() {
	suite.T().Log("Running TransformUnorderedMapSuite: Tests for ordered_union")

	even := r.DB("test").Table("even")
	_ = even // Prevent any noused variable errors
	odd := r.DB("test").Table("odd")
	_ = odd // Prevent any noused variable errors
	odd2 := r.DB("test").Table("odd2")
	_ = odd2 // Prevent any noused variable errors


	{
		// transform/unordered_map.yaml line #6
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* odd.insert([{"id":1, "num":1}, {"id":3, "num":3}, {"id":5, "num":5}]) */

		suite.T().Log("About to run line #6: odd.Insert([]interface{}{map[interface{}]interface{}{'id': 1, 'num': 1, }, map[interface{}]interface{}{'id': 3, 'num': 3, }, map[interface{}]interface{}{'id': 5, 'num': 5, }})")

		runAndAssert(suite.Suite, expected_, odd.Insert([]interface{}{map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 5, "num": 5, }}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// transform/unordered_map.yaml line #7
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* even.insert([{"id":2, "num":2}, {"id":4, "num":4}, {"id":6, "num":6}]) */

		suite.T().Log("About to run line #7: even.Insert([]interface{}{map[interface{}]interface{}{'id': 2, 'num': 2, }, map[interface{}]interface{}{'id': 4, 'num': 4, }, map[interface{}]interface{}{'id': 6, 'num': 6, }})")

		runAndAssert(suite.Suite, expected_, even.Insert([]interface{}{map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 6, "num": 6, }}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// transform/unordered_map.yaml line #8
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* odd2.insert([{"id":7, "num":1}, {"id":8, "num":3}, {"id":9, "num":2}]) */

		suite.T().Log("About to run line #8: odd2.Insert([]interface{}{map[interface{}]interface{}{'id': 7, 'num': 1, }, map[interface{}]interface{}{'id': 8, 'num': 3, }, map[interface{}]interface{}{'id': 9, 'num': 2, }})")

		runAndAssert(suite.Suite, expected_, odd2.Insert([]interface{}{map[interface{}]interface{}{"id": 7, "num": 1, }, map[interface{}]interface{}{"id": 8, "num": 3, }, map[interface{}]interface{}{"id": 9, "num": 2, }}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// transform/unordered_map.yaml line #11
		/* [{"id":1, "num":1}, {"id":3, "num":3}, {"id":5, "num":5}, {"id":2, "num":2}, {"id":4, "num":4}, {"id":6, "num":6}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 5, "num": 5, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 6, "num": 6, }}
		/* odd.order_by("num").union(even.order_by("num"), interleave = false) */

		suite.T().Log("About to run line #11: odd.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: false, }, even.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: false, }, even.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// transform/unordered_map.yaml line #16
		/* [{"id":2, "num":2}, {"id":4, "num":4}, {"id":6, "num":6}, {"id":1, "num":1}, {"id":3, "num":3}, {"id":5, "num":5}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 6, "num": 6, }, map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 5, "num": 5, }}
		/* even.order_by("num").union(odd.order_by("num"), interleave = false) */

		suite.T().Log("About to run line #16: even.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: false, }, odd.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, even.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: false, }, odd.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// transform/unordered_map.yaml line #22
		/* [{"id":1, "num":1}, {"id":2, "num":2}, {"id":3, "num":3}, {"id":4, "num":4}, {"id":5, "num":5}, {"id":6, "num":6}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 5, "num": 5, }, map[interface{}]interface{}{"id": 6, "num": 6, }}
		/* odd.order_by("num").union(even.order_by("num"), interleave="num") */

		suite.T().Log("About to run line #22: odd.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: 'num', }, even.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: "num", }, even.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// transform/unordered_map.yaml line #28
		/* err("ReqlQueryLogicError","The streams given as arguments are not ordered by given ordering.") */
		var expected_ Err = err("ReqlQueryLogicError", "The streams given as arguments are not ordered by given ordering.")
		/* odd.order_by("num").union(even.order_by("num"), interleave=r.desc("num")) */

		suite.T().Log("About to run line #28: odd.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: r.Desc('num'), }, even.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: r.Desc("num"), }, even.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// transform/unordered_map.yaml line #34
		/* [{"id":1, "num":1}, {"id":2, "num":2}, {"id":3, "num":3}, {"id":4, "num":4}, {"id":5, "num":5}, {"id":6, "num":6}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 5, "num": 5, }, map[interface{}]interface{}{"id": 6, "num": 6, }}
		/* odd.order_by("num").union(even.order_by("num"), interleave=lambda x: x["num"]) */

		suite.T().Log("About to run line #34: odd.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: func(x r.Term) interface{} { return x.AtIndex('num')}, }, even.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: func(x r.Term) interface{} { return x.AtIndex("num")}, }, even.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// transform/unordered_map.yaml line #40
		/* [{"id": 7, "num": 1}, {"id": 2, "num": 2}, {"id": 9, "num": 2}, {"id": 8, "num": 3}, {"id": 4, "num": 4}, {"id": 6, "num": 6}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 7, "num": 1, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 9, "num": 2, }, map[interface{}]interface{}{"id": 8, "num": 3, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 6, "num": 6, }}
		/* odd2.order_by("num", r.desc("id")).union(even.order_by("num", r.desc("id")), interleave=[lambda x: x["num"], lambda x: x["id"]]) */

		suite.T().Log("About to run line #40: odd2.OrderBy('num', r.Desc('id')).UnionWithOpts(r.UnionOpts{Interleave: []interface{}{func(x r.Term) interface{} { return x.AtIndex('num')}, func(x r.Term) interface{} { return x.AtIndex('id')}}, }, even.OrderBy('num', r.Desc('id')))")

		runAndAssert(suite.Suite, expected_, odd2.OrderBy("num", r.Desc("id")).UnionWithOpts(r.UnionOpts{Interleave: []interface{}{func(x r.Term) interface{} { return x.AtIndex("num")}, func(x r.Term) interface{} { return x.AtIndex("id")}}, }, even.OrderBy("num", r.Desc("id"))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// transform/unordered_map.yaml line #46
		/* err("ReqlServerCompileError", "DESC may only be used as an argument to ORDER_BY or UNION.") */
		var expected_ Err = err("ReqlCompileError", "DESC may only be used as an argument to ORDER_BY or UNION.")
		/* odd.order_by("num").union(even.order_by("num"), interleave=lambda x: r.desc(x["num"])) */

		suite.T().Log("About to run line #46: odd.OrderBy('num').UnionWithOpts(r.UnionOpts{Interleave: func(x r.Term) interface{} { return r.Desc(x.AtIndex('num'))}, }, even.OrderBy('num'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num").UnionWithOpts(r.UnionOpts{Interleave: func(x r.Term) interface{} { return r.Desc(x.AtIndex("num"))}, }, even.OrderBy("num")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// transform/unordered_map.yaml line #50
		/* [{"id":6, "num":6}, {"id":5, "num":5}, {"id":4, "num":4}, {"id":3, "num":3}, {"id":2, "num":2}, {"id":1, "num":1}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 6, "num": 6, }, map[interface{}]interface{}{"id": 5, "num": 5, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 1, "num": 1, }}
		/* odd.order_by(r.desc("num")).union(even.order_by(r.desc("num")), interleave= [r.desc(lambda x: x["num"])]) */

		suite.T().Log("About to run line #50: odd.OrderBy(r.Desc('num')).UnionWithOpts(r.UnionOpts{Interleave: []interface{}{r.Desc(func(x r.Term) interface{} { return x.AtIndex('num')})}, }, even.OrderBy(r.Desc('num')))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy(r.Desc("num")).UnionWithOpts(r.UnionOpts{Interleave: []interface{}{r.Desc(func(x r.Term) interface{} { return x.AtIndex("num")})}, }, even.OrderBy(r.Desc("num"))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// transform/unordered_map.yaml line #54
		/* [{"id":1, "num":1}, {"id":7, "num":1}, {"id":2, "num":2}, {"id":9, "num":2}, {"id":3, "num":3}, {"id":8, "num":3}, {"id":4, "num":4}, {"id":5, "num":5}, {"id":6, "num":6}, ] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, "num": 1, }, map[interface{}]interface{}{"id": 7, "num": 1, }, map[interface{}]interface{}{"id": 2, "num": 2, }, map[interface{}]interface{}{"id": 9, "num": 2, }, map[interface{}]interface{}{"id": 3, "num": 3, }, map[interface{}]interface{}{"id": 8, "num": 3, }, map[interface{}]interface{}{"id": 4, "num": 4, }, map[interface{}]interface{}{"id": 5, "num": 5, }, map[interface{}]interface{}{"id": 6, "num": 6, }}
		/* odd.order_by("num", "id").union(even.order_by("num", "id"), odd2.order_by("num", "id"), interleave= ["num", "id"]) */

		suite.T().Log("About to run line #54: odd.OrderBy('num', 'id').UnionWithOpts(r.UnionOpts{Interleave: []interface{}{'num', 'id'}, }, even.OrderBy('num', 'id'), odd2.OrderBy('num', 'id'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num", "id").UnionWithOpts(r.UnionOpts{Interleave: []interface{}{"num", "id"}, }, even.OrderBy("num", "id"), odd2.OrderBy("num", "id")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// transform/unordered_map.yaml line #58
		/* err("ReqlQueryLogicError","The streams given as arguments are not ordered by given ordering.") */
		var expected_ Err = err("ReqlQueryLogicError", "The streams given as arguments are not ordered by given ordering.")
		/* odd.order_by("num", "id").union(even.order_by("num", "id"), odd2.order_by(r.desc("num"), "id"), interleave= ["num", "id"]) */

		suite.T().Log("About to run line #58: odd.OrderBy('num', 'id').UnionWithOpts(r.UnionOpts{Interleave: []interface{}{'num', 'id'}, }, even.OrderBy('num', 'id'), odd2.OrderBy(r.Desc('num'), 'id'))")

		runAndAssert(suite.Suite, expected_, odd.OrderBy("num", "id").UnionWithOpts(r.UnionOpts{Interleave: []interface{}{"num", "id"}, }, even.OrderBy("num", "id"), odd2.OrderBy(r.Desc("num"), "id")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #58")
	}
}
