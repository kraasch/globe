package geocalc

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	// other imports.
	util "github.com/kraasch/geo/pkg/testutil"
	godiff "github.com/kraasch/godiff/godiff"
)

const (
	dateLayout = "2006-01-02" // Go's reference date layout.
)

var NL = fmt.Sprintln()

type TestList struct {
	testName      string
	isMulti       bool
	inputArr      []string
	expectedValue string
}

type TestSuite struct {
	testingFunction func(in TestList) string
	tests           []TestList
}

var suites = []TestSuite{ // All tests.

	/*
	 * Test for the function MoonPhase().
	 */
	{
		testingFunction: func(in TestList) string {
			time, err0 := time.Parse(dateLayout, in.inputArr[0])
			if err0 != nil {
				return "error in type converstion within the test: date."
			}
			out := MoonPhase(time)
			return out
		},
		tests: []TestList{
			{
				testName: "sun_moon-phase_non-verbose_calculate_00",
				isMulti:  true,
				inputArr: []string{
					"2000-01-01", // Some date.
				},
				expectedValue: //
				" ● phase:   Last Quarter" + NL +
					" ● age:     24.38 days (◖)" + NL +
					" ● dist.:   398596 km" + NL +
					" ● illum.:  27%" + NL +
					" ● new in:  5.8 days" + NL +
					" ● full in: 20.2 days" + NL +
					" ● new on:  2000-01-06, Thu" + NL +
					" ● full on: 2000-01-21, Fri",
			},
		},
	},

	/*
	 * Test for the function MoonPhaseVerbose().
	 */
	{
		testingFunction: func(in TestList) string {
			time, err0 := time.Parse(dateLayout, in.inputArr[0])
			if err0 != nil {
				return "error in type converstion within the test: date."
			}
			out := MoonPhaseVerbose(time)
			return out
		},
		tests: []TestList{
			{
				testName: "sun_moon-phase_verbose_calculate_00",
				isMulti:  true,
				inputArr: []string{
					"2000-01-01", // Some date.
				},
				expectedValue: //
				"The moon is 24.38 days old, and is therefore in Last Quarter phase (◖)." + NL +
					"It is 398596 km from the centre of the Earth." + NL +
					// TODO: xxx percent of distance between 363104 (min) and 405500 (max).
					"It is 27% illuminated." + NL +
					"The next new moon is in 5.8 days (2000-01-06, Thu)." + NL +
					"The next full moon is in 20.2 days (2000-01-21, Fri).",
				/* NOTE: other values are:
				 *  - PHASE:     0.8255741837703208
				 *  - ILLUM:     0.27139898737765766
				 *  - AGE:       24.379691645748075
				 *  - DIST:      398596.29455439356
				 *  - angdia:    0.49964879458462286
				 *  - sundist:   1.4710022336390877e+08
				 *  - sunangdia: 0.5421823793611453
				 *  - pdata:     2.4515445e+06
				 *  - quarters:  [8]float64{9.446059836222186e+08, 9.453054569995925e+08, 9.458839879086032e+08, 9.46476358262685e+08, 9.471824973697052e+08, 9.478568800435827e+08, 9.484297056796163e+08, 9.490463221789911e+08}
				 *  - timespace: 9.466848e+08
				 *  - LONGITUDE: 217.09009941697522
				 */
			},
		},
	},

	/*
	 * Test for the function SunRiseAndSet().
	 */
	{
		testingFunction: func(in TestList) string {
			lat, err0 := strconv.ParseFloat(in.inputArr[0], 64)
			lon, err1 := strconv.ParseFloat(in.inputArr[1], 64)
			date, err2 := time.Parse(dateLayout, in.inputArr[2])
			if err0 != nil {
				return "error in type converstion within the test: first float."
			}
			if err1 != nil {
				return "error in type converstion within the test: second float."
			}
			if err2 != nil {
				return "error in type converstion within the test: date."
			}
			out := SunRiseAndSet(lat, lon, date)
			return out
		},
		tests: []TestList{
			{
				testName: "sun_sunrise+sunset_calculate_00",
				isMulti:  true,
				inputArr: []string{
					"43.65", "-79.38", // Toronto, Canada.
					"2000-01-01", // Some date.
				},
				// expectedValue: "sunrise: 12:51, sunset: 21:51", // NOTE: as UTC.
				expectedValue: // NOTE: as GMT-5 (in Toronto, Canada).
				" ☼ rise:    07:51 h" + NL +
					" ☼ set:     16:51 h",
			},
		},
	},

	/*
	 * Test for the function LocalAndUtcTime().
	 */
	{
		testingFunction: func(in TestList) string {
			t := LocalAndUtcTime()
			return util.AnonymizeTimeStrings(t)
		},
		tests: []TestList{
			{
				testName: "sun_sunrise+sunset_calculate_00",
				isMulti:  true,
				inputArr: []string{},
				expectedValue: // line break.
				util.AnonymizeTimeStrings(" ▣ time:    17:00 h") + NL +
					util.AnonymizeTimeStrings(" 🜃 utc:     03:12 h"),
			},
		},
	},
}

func TestAll(t *testing.T) {
	for _, suite := range suites {
		for _, test := range suite.tests {
			name := test.testName
			t.Run(name, func(t *testing.T) {
				exp := test.expectedValue
				got := suite.testingFunction(test)
				if exp != got {
					if test.isMulti {
						t.Errorf("In '%s':\n", name)
						diff := godiff.CDiff(exp, got)
						t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
						t.Errorf("exp/got:\n%s\n", diff)
					} else {
						t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
					}
				}
			})
		}
	}
}
