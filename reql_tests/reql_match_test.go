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
func TestMatchSuite(t *testing.T) {
    suite.Run(t, new(MatchSuite ))
}

type MatchSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MatchSuite) SetupTest() {
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

func (suite *MatchSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
	 r.DB("test").TableDrop("tbl").Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *MatchSuite) TestCases() {
	tbl := r.DB("test").Table("tbl")


    {
        // match.yaml line #4
        /* ({'str':'bcde','groups':[null,{'start':2,'str':'cde','end':5}],'start':1,'end':5}) */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"str": "bcde", "groups": []interface{}{nil, map[interface{}]interface{}{"start": 2, "str": "cde", "end": 5, }}, "start": 1, "end": 5, }
        /* r.expr("abcdefg").match("a(b.e)|b(c.e)") */

    	suite.T().Log("About to run line #4: r.Expr('abcdefg').Match('a(b.e)|b(c.e)')")

        cursor, err := r.Expr("abcdefg").Match("a(b.e)|b(c.e)").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #4")
    }

    {
        // match.yaml line #6
        /* (null) */
        var expected_ interface{} = nil
        /* r.expr("abcdefg").match("a(b.e)|B(c.e)") */

    	suite.T().Log("About to run line #6: r.Expr('abcdefg').Match('a(b.e)|B(c.e)')")

        cursor, err := r.Expr("abcdefg").Match("a(b.e)|B(c.e)").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #6")
    }

    {
        // match.yaml line #8
        /* ({'str':'bcde','groups':[null,{'start':2,'str':'cde','end':5}],'start':1,'end':5}) */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"str": "bcde", "groups": []interface{}{nil, map[interface{}]interface{}{"start": 2, "str": "cde", "end": 5, }}, "start": 1, "end": 5, }
        /* r.expr("abcdefg").match("(?i)a(b.e)|B(c.e)") */

    	suite.T().Log("About to run line #8: r.Expr('abcdefg').Match('(?i)a(b.e)|B(c.e)')")

        cursor, err := r.Expr("abcdefg").Match("(?i)a(b.e)|B(c.e)").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #8")
    }

    {
        // match.yaml line #12
        /* (["aca", "ada"]) */
        var expected_ []interface{} = []interface{}{"aca", "ada"}
        /* r.expr(["aba", "aca", "ada", "aea"]).filter(lambda row:row.match("a(.)a")['groups'][0]['str'].match("[cd]")) */

    	suite.T().Log("About to run line #12: r.Expr([]interface{}{'aba', 'aca', 'ada', 'aea'}).Filter(func(row r.Term) interface{} { return row.Match('a(.)a').AtIndex('groups').AtIndex(0).AtIndex('str').Match('[cd]')})")

        cursor, err := r.Expr([]interface{}{"aba", "aca", "ada", "aea"}).Filter(func(row r.Term) interface{} { return row.Match("a(.)a").AtIndex("groups").AtIndex(0).AtIndex("str").Match("[cd]")}).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #12")
    }

    {
        // match.yaml line #16
        /* ({'deleted':0,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':3}) */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 3, }
        /* tbl.insert([{'id':0,'a':'abc'},{'id':1,'a':'ab'},{'id':2,'a':'bc'}]) */

    	suite.T().Log("About to run line #16: tbl.Insert([]interface{}{map[interface{}]interface{}{'id': 0, 'a': 'abc', }, map[interface{}]interface{}{'id': 1, 'a': 'ab', }, map[interface{}]interface{}{'id': 2, 'a': 'bc', }})")

        cursor, err := tbl.Insert([]interface{}{map[interface{}]interface{}{"id": 0, "a": "abc", }, map[interface{}]interface{}{"id": 1, "a": "ab", }, map[interface{}]interface{}{"id": 2, "a": "bc", }}).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #16")
    }

    {
        // match.yaml line #20
        /* ([{'id':0,'a':'abc'},{'id':1,'a':'ab'},{'id':2,'a':'bc'}]) */
        var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 0, "a": "abc", }, map[interface{}]interface{}{"id": 1, "a": "ab", }, map[interface{}]interface{}{"id": 2, "a": "bc", }}
        /* tbl.filter(lambda row:row['a'].match('b')).order_by('id') */

    	suite.T().Log("About to run line #20: tbl.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Match('b')}).OrderBy('id')")

        cursor, err := tbl.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Match("b")}).OrderBy("id").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #20")
    }

    {
        // match.yaml line #24
        /* ([{'id':0,'a':'abc'},{'id':1,'a':'ab'}]) */
        var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 0, "a": "abc", }, map[interface{}]interface{}{"id": 1, "a": "ab", }}
        /* tbl.filter(lambda row:row['a'].match('ab')).order_by('id') */

    	suite.T().Log("About to run line #24: tbl.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Match('ab')}).OrderBy('id')")

        cursor, err := tbl.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Match("ab")}).OrderBy("id").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #24")
    }

    {
        // match.yaml line #28
        /* ([{'id':1,'a':'ab'}]) */
        var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, "a": "ab", }}
        /* tbl.filter(lambda row:row['a'].match('ab$')).order_by('id') */

    	suite.T().Log("About to run line #28: tbl.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Match('ab$')}).OrderBy('id')")

        cursor, err := tbl.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Match("ab$")}).OrderBy("id").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #28")
    }

    {
        // match.yaml line #32
        /* ([]) */
        var expected_ []interface{} = []interface{}{}
        /* tbl.filter(lambda row:row['a'].match('^b$')).order_by('id') */

    	suite.T().Log("About to run line #32: tbl.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Match('^b$')}).OrderBy('id')")

        cursor, err := tbl.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Match("^b$")}).OrderBy("id").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #32")
    }

    {
        // match.yaml line #36
        /* err("ReqlQueryLogicError", "Error in regexp `ab\\9` (portion `\\9`): invalid escape sequence: \\9", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Error in regexp `ab\\9` (portion `\\9`): invalid escape sequence: \\9")
        /* r.expr("").match("ab\\9") */

    	suite.T().Log("About to run line #36: r.Expr('').Match('ab\\\\9')")

        cursor, err := r.Expr("").Match("ab\\9").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #36")
    }
}