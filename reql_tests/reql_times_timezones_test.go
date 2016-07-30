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

// Test basic timezone manipulation
func TestTimesTimezonesSuite(t *testing.T) {
	suite.Run(t, new(TimesTimezonesSuite ))
}

type TimesTimezonesSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesTimezonesSuite) SetupTest() {
	suite.T().Log("Setting up TimesTimezonesSuite")
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

func (suite *TimesTimezonesSuite) TearDownSuite() {
	suite.T().Log("Tearing down TimesTimezonesSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TimesTimezonesSuite) TestCases() {
	suite.T().Log("Running TimesTimezonesSuite: Test basic timezone manipulation")



	// times/timezones.yaml line #3
	// t1 = r.time(2013, r.july, 29, 23, 30, 0, "+00:00")
	suite.T().Log("Possibly executing: var t1 r.Term = r.Time(2013, r.July, 29, 23, 30, 0, '+00:00')")

	t1 := r.Time(2013, r.July, 29, 23, 30, 0, "+00:00")
	_ = t1 // Prevent any noused variable errors


	// times/timezones.yaml line #5
	// tutc1 = t1.in_timezone("Z")
	suite.T().Log("Possibly executing: var tutc1 r.Term = t1.InTimezone('Z')")

	tutc1 := t1.InTimezone("Z")
	_ = tutc1 // Prevent any noused variable errors


	// times/timezones.yaml line #6
	// tutc2 = t1.in_timezone("+00:00")
	suite.T().Log("Possibly executing: var tutc2 r.Term = t1.InTimezone('+00:00')")

	tutc2 := t1.InTimezone("+00:00")
	_ = tutc2 // Prevent any noused variable errors


	// times/timezones.yaml line #7
	// tutc3 = t1.in_timezone("+00")
	suite.T().Log("Possibly executing: var tutc3 r.Term = t1.InTimezone('+00')")

	tutc3 := t1.InTimezone("+00")
	_ = tutc3 // Prevent any noused variable errors


	// times/timezones.yaml line #8
	// tutcs = r.expr([tutc1, tutc2, tutc3])
	suite.T().Log("Possibly executing: var tutcs r.Term = r.Expr([]interface{}{tutc1, tutc2, tutc3})")

	tutcs := r.Expr([]interface{}{tutc1, tutc2, tutc3})
	_ = tutcs // Prevent any noused variable errors


	// times/timezones.yaml line #10
	// tm1 = t1.in_timezone("-00:59")
	suite.T().Log("Possibly executing: var tm1 r.Term = t1.InTimezone('-00:59')")

	tm1 := t1.InTimezone("-00:59")
	_ = tm1 // Prevent any noused variable errors


	// times/timezones.yaml line #11
	// tm2 = t1.in_timezone("-01:00")
	suite.T().Log("Possibly executing: var tm2 r.Term = t1.InTimezone('-01:00')")

	tm2 := t1.InTimezone("-01:00")
	_ = tm2 // Prevent any noused variable errors


	// times/timezones.yaml line #12
	// tm3 = t1.in_timezone("-01:01")
	suite.T().Log("Possibly executing: var tm3 r.Term = t1.InTimezone('-01:01')")

	tm3 := t1.InTimezone("-01:01")
	_ = tm3 // Prevent any noused variable errors


	// times/timezones.yaml line #13
	// tms = r.expr([tm1, tm2, tm3])
	suite.T().Log("Possibly executing: var tms r.Term = r.Expr([]interface{}{tm1, tm2, tm3})")

	tms := r.Expr([]interface{}{tm1, tm2, tm3})
	_ = tms // Prevent any noused variable errors


	// times/timezones.yaml line #15
	// tp1 = t1.in_timezone("+00:59")
	suite.T().Log("Possibly executing: var tp1 r.Term = t1.InTimezone('+00:59')")

	tp1 := t1.InTimezone("+00:59")
	_ = tp1 // Prevent any noused variable errors


	// times/timezones.yaml line #16
	// tp2 = t1.in_timezone("+01:00")
	suite.T().Log("Possibly executing: var tp2 r.Term = t1.InTimezone('+01:00')")

	tp2 := t1.InTimezone("+01:00")
	_ = tp2 // Prevent any noused variable errors


	// times/timezones.yaml line #17
	// tp3 = t1.in_timezone("+01:01")
	suite.T().Log("Possibly executing: var tp3 r.Term = t1.InTimezone('+01:01')")

	tp3 := t1.InTimezone("+01:01")
	_ = tp3 // Prevent any noused variable errors


	// times/timezones.yaml line #18
	// tps = r.expr([tp1, tp2, tp3])
	suite.T().Log("Possibly executing: var tps r.Term = r.Expr([]interface{}{tp1, tp2, tp3})")

	tps := r.Expr([]interface{}{tp1, tp2, tp3})
	_ = tps // Prevent any noused variable errors


	// times/timezones.yaml line #20
	// ts = tutcs.union(tms).union(tps).union([t1])
	suite.T().Log("Possibly executing: var ts r.Term = tutcs.Union(tms).Union(tps).Union([]interface{}{t1})")

	ts := tutcs.Union(tms).Union(tps).Union([]interface{}{t1})
	_ = ts // Prevent any noused variable errors


	{
		// times/timezones.yaml line #23
		/* ([["+00:00", 29], ["+00:00", 29], ["+00:00", 29]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{"+00:00", 29}, []interface{}{"+00:00", 29}, []interface{}{"+00:00", 29}}
		/* tutcs.map(lambda x:[x.timezone(), x.day()]) */

		suite.T().Log("About to run line #23: tutcs.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}})")

		runAndAssert(suite.Suite, expected_, tutcs.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// times/timezones.yaml line #27
		/* ([["-00:59", 29], ["-01:00", 29], ["-01:01", 29]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{"-00:59", 29}, []interface{}{"-01:00", 29}, []interface{}{"-01:01", 29}}
		/* tms.map(lambda x:[x.timezone(), x.day()]) */

		suite.T().Log("About to run line #27: tms.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}})")

		runAndAssert(suite.Suite, expected_, tms.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// times/timezones.yaml line #31
		/* ([["+00:59", 30], ["+01:00", 30], ["+01:01", 30]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{"+00:59", 30}, []interface{}{"+01:00", 30}, []interface{}{"+01:01", 30}}
		/* tps.map(lambda x:[x.timezone(), x.day()]) */

		suite.T().Log("About to run line #31: tps.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}})")

		runAndAssert(suite.Suite, expected_, tps.Map(func(x r.Term) interface{} { return []interface{}{x.Timezone(), x.Day()}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// times/timezones.yaml line #38
		/* ([0]) */
		var expected_ []interface{} = []interface{}{0}
		/* ts.concat_map(lambda x:ts.map(lambda y:x - y)).distinct() */

		suite.T().Log("About to run line #38: ts.ConcatMap(func(x r.Term) interface{} { return ts.Map(func(y r.Term) interface{} { return r.Sub(x, y)})}).Distinct()")

		runAndAssert(suite.Suite, expected_, ts.ConcatMap(func(x r.Term) interface{} { return ts.Map(func(y r.Term) interface{} { return r.Sub(x, y)})}).Distinct(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// times/timezones.yaml line #44
		/* err('ReqlQueryLogicError', 'Timezone `` does not start with `-` or `+`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Timezone `` does not start with `-` or `+`.")
		/* r.now().in_timezone("") */

		suite.T().Log("About to run line #44: r.Now().InTimezone('')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone(""), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// times/timezones.yaml line #47
		/* err('ReqlQueryLogicError', '`-00` is not a valid time offset.') */
		var expected_ Err = err("ReqlQueryLogicError", "`-00` is not a valid time offset.")
		/* r.now().in_timezone("-00") */

		suite.T().Log("About to run line #47: r.Now().InTimezone('-00')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone("-00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #47")
	}

	{
		// times/timezones.yaml line #50
		/* err('ReqlQueryLogicError', '`-00:00` is not a valid time offset.') */
		var expected_ Err = err("ReqlQueryLogicError", "`-00:00` is not a valid time offset.")
		/* r.now().in_timezone("-00:00") */

		suite.T().Log("About to run line #50: r.Now().InTimezone('-00:00')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone("-00:00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// times/timezones.yaml line #53
		/* err('ReqlQueryLogicError', 'Timezone `UTC+00` does not start with `-` or `+`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Timezone `UTC+00` does not start with `-` or `+`.")
		/* r.now().in_timezone("UTC+00") */

		suite.T().Log("About to run line #53: r.Now().InTimezone('UTC+00')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone("UTC+00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// times/timezones.yaml line #56
		/* err('ReqlQueryLogicError', 'Minutes out of range in `+00:60`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Minutes out of range in `+00:60`.")
		/* r.now().in_timezone("+00:60") */

		suite.T().Log("About to run line #56: r.Now().InTimezone('+00:60')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone("+00:60"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #56")
	}

	{
		// times/timezones.yaml line #59
		/* err('ReqlQueryLogicError', 'Hours out of range in `+25:00`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Hours out of range in `+25:00`.")
		/* r.now().in_timezone("+25:00") */

		suite.T().Log("About to run line #59: r.Now().InTimezone('+25:00')")

		runAndAssert(suite.Suite, expected_, r.Now().InTimezone("+25:00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #59")
	}

	{
		// times/timezones.yaml line #63
		/* err('ReqlQueryLogicError', 'Timezone `` does not start with `-` or `+`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Timezone `` does not start with `-` or `+`.")
		/* r.time(2013, 1, 1, "") */

		suite.T().Log("About to run line #63: r.Time(2013, 1, 1, '')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, ""), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #63")
	}

	{
		// times/timezones.yaml line #66
		/* err('ReqlQueryLogicError', '`-00` is not a valid time offset.') */
		var expected_ Err = err("ReqlQueryLogicError", "`-00` is not a valid time offset.")
		/* r.time(2013, 1, 1, "-00") */

		suite.T().Log("About to run line #66: r.Time(2013, 1, 1, '-00')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, "-00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// times/timezones.yaml line #69
		/* err('ReqlQueryLogicError', '`-00:00` is not a valid time offset.') */
		var expected_ Err = err("ReqlQueryLogicError", "`-00:00` is not a valid time offset.")
		/* r.time(2013, 1, 1, "-00:00") */

		suite.T().Log("About to run line #69: r.Time(2013, 1, 1, '-00:00')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, "-00:00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #69")
	}

	{
		// times/timezones.yaml line #72
		/* err('ReqlQueryLogicError', 'Timezone `UTC+00` does not start with `-` or `+`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Timezone `UTC+00` does not start with `-` or `+`.")
		/* r.time(2013, 1, 1, "UTC+00") */

		suite.T().Log("About to run line #72: r.Time(2013, 1, 1, 'UTC+00')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, "UTC+00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #72")
	}

	{
		// times/timezones.yaml line #75
		/* err('ReqlQueryLogicError', 'Minutes out of range in `+00:60`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Minutes out of range in `+00:60`.")
		/* r.time(2013, 1, 1, "+00:60") */

		suite.T().Log("About to run line #75: r.Time(2013, 1, 1, '+00:60')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, "+00:60"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #75")
	}

	{
		// times/timezones.yaml line #78
		/* err('ReqlQueryLogicError', 'Hours out of range in `+25:00`.') */
		var expected_ Err = err("ReqlQueryLogicError", "Hours out of range in `+25:00`.")
		/* r.time(2013, 1, 1, "+25:00") */

		suite.T().Log("About to run line #78: r.Time(2013, 1, 1, '+25:00')")

		runAndAssert(suite.Suite, expected_, r.Time(2013, 1, 1, "+25:00"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #78")
	}

	{
		// times/timezones.yaml line #81
		/* ("2015-07-08T00:00:00-08:00") */
		var expected_ string = "2015-07-08T00:00:00-08:00"
		/* r.epoch_time(1436428422.339).in_timezone('-08:00').date().to_iso8601() */

		suite.T().Log("About to run line #81: r.EpochTime(1436428422.339).InTimezone('-08:00').Date().ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1436428422.339).InTimezone("-08:00").Date().ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// times/timezones.yaml line #85
		/* ("2015-07-09T00:00:00-07:00") */
		var expected_ string = "2015-07-09T00:00:00-07:00"
		/* r.epoch_time(1436428422.339).in_timezone('-07:00').date().to_iso8601() */

		suite.T().Log("About to run line #85: r.EpochTime(1436428422.339).InTimezone('-07:00').Date().ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1436428422.339).InTimezone("-07:00").Date().ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #85")
	}
}
