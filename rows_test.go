package rethinkgo

import (
	"fmt"
	test "launchpad.net/gocheck"
	"testing"
)

type object struct {
	Id    int64  `rethinkdb:"id"`
	Name  string `rethinkdb:"name"`
	Attrs []attr
}

type attr struct {
	Name  string
	Value interface{}
}

func (s *RethinkSuite) TestResultScanLiteral(c *test.C) {
	rows, err := Expr(5).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response interface{}
		rows.Scan(&response)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanSlice(c *test.C) {
	rows, err := Expr(List{1, 2, 3, 4, 5}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response interface{}
		rows.Scan(&response)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanIntSlice(c *test.C) {
	rows, err := Expr(List{1, 2, 3, 4, 5}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response []int
		rows.Scan(&response)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanMap(c *test.C) {
	rows, err := Expr(Obj{
		"id":   2,
		"name": "Object 1",
		"attr": List{Obj{
			"name":  "attr 1",
			"value": "value 1",
		}},
	}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response map[string]interface{}
		err = rows.Scan(&response)
		c.Assert(err, test.IsNil)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanMapIntoInterface(c *test.C) {
	rows, err := Expr(Obj{
		"id":   2,
		"name": "Object 1",
		"attr": List{Obj{
			"name":  "attr 1",
			"value": "value 1",
		}},
	}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response interface{}
		err = rows.Scan(&response)
		c.Assert(err, test.IsNil)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanMapNested(c *test.C) {
	rows, err := Expr(Obj{
		"id":   2,
		"name": "Object 1",
		"attr": List{Obj{
			"name":  "attr 1",
			"value": "value 1",
		}},
	}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response map[string]interface{}
		err = rows.Scan(&response)
		c.Assert(err, test.IsNil)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultScanStruct(c *test.C) {
	rows, err := Expr(Obj{
		"id":   2,
		"name": "Object 1",
		"Attrs": List{Obj{
			"Name":  "attr 1",
			"Value": "value 1",
		}},
	}).Run(conn)
	c.Assert(err, test.IsNil)

	for rows.Next() {
		var response object
		err = rows.Scan(&response)
		c.Assert(err, test.IsNil)

		fmt.Printf("%#v\n", response)
	}
}

func (s *RethinkSuite) TestResultAtomString(c *test.C) {
	// row := Expr("a").RunRow(conn)
	// c.Assert(row, test.Equals, "a")
}

func (s *RethinkSuite) TestResultAtomArray(c *test.C) {
	// query := Expr(List{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	// result, err := query.Run(conn)
	// c.Assert(err, test.IsNil)

	// num := 0
	// for result.Next() {
	// 	row, err := result.Row()
	// 	c.Assert(err, test.IsNil)
	// 	c.Assert(len(row.([]interface{})), test.Equals, 10)
	// 	num++
	// }

	// c.Assert(num, test.Equals, 1)
}

func (s *RethinkSuite) TestResultPartial(c *test.C) {
	// Since this test takes a long time to run skip if the short flag is set
	if testing.Short() {
		c.Skip("-test.short is set")
	}

	// Ensure table is empty to start with
	_, err := Db("test").Table("test").Delete().Run(conn)
	c.Assert(err, test.IsNil)

	// Insert 1500 test rows
	result, err := Js("var vars = []; for (var i=0; i<1500; i++) {vars.push(i);} vars").ForEach(func(i RqlTerm) RqlTerm {
		return Db("test").Table("test").Insert(Obj{"num": i})
	}).Run(conn)
	c.Assert(err, test.IsNil)

	// Ensure that 1500 rows are returned
	result, err = Db("test").Table("test").Run(conn)
	c.Assert(err, test.IsNil)

	num := 0
	for result.Next() {
		num++
	}

	c.Assert(num, test.Equals, 1500)

	// Delete test rows
	_, err = Db("test").Table("test").Delete().Run(conn)
	c.Assert(err, test.IsNil)
}