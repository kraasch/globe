package geomap

import (

	// this is a test.
	"strconv"
	"testing"

	// printing and formatting.

	// other imports.
	"github.com/kraasch/godiff/godiff"
)

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
	 * Test for the function PrintCoord(lon, lat).
	 */
	{
		testingFunction: func(in TestList) string {
			lat, err0 := strconv.ParseFloat(in.inputArr[0], 64)
			lon, err1 := strconv.ParseFloat(in.inputArr[1], 64)
			if err0 != nil {
				return "error in type converstion within the test: first float."
			}
			if err1 != nil {
				return "error in type converstion within the test: second float."
			}
			world := NewWorld()
			out, reserr := world.PrintCoord(lat, lon)
			if reserr != nil {
				return "error getting a result."
			}
			return out
		},
		tests: []TestList{
			{
				testName: "map_pretty-print_coord_00",
				isMulti:  false,
				inputArr: []string{
					// mark the following location with '▣'.
					"0.0",    // latitude, ie (=).
					"-180.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				//"1-987654321 123456789+12"
				"    _,--._  _._.--.--.._" + NL + //   lat. +50 to +90. = 40 degrees.  + ==> North.
					"=.--'=_',-,:`;_      .,'" + NL + // lat. +50 to +35. = 15 degrees.  - ==> South.
					",-.  _.)  (``-;_   .'   " + NL + // lat. +35 to +20. = 15 degrees.
					"   '-:_    `) ) .''=.   " + NL + // lat. +20 to -20. = 40 degrees.
					"     ) )    ()'    ='   " + NL + // lat. -20 to -35. = 15 degrees.
					"     |/            (_) =" + NL + // lat. -35 to -50. = 15 degrees.
					"     -                  ", //       lat. -50 to -90. = 40 degrees.
				// A          A           A
				// |          |           |
				// |          0           | ==> 360/24 = 15        // 24 columns, 360 degrees on the globe.
				// |                      |                        // thus: 0-15 degrees are the first column, and so on.
				// |                      | func lon2col( lon ):   // write a function to translate latitude to columns.
				// .__ -180(E)  +180(W) __.   lon = lon+180        // start with 0 at the left.
				//                            return int(lon / 15) // divide without rest.
			},
		},
	}, // End of all tests.

	/*
	 * Test for the function PrintInner().
	 */
	{
		testingFunction: func(in TestList) string {
			world := NewWorld()
			out := world.PrintInner()
			return out
		},
		tests: []TestList{
			{
				testName: "map_pretty-print_inner_00",
				isMulti:  false,
				inputArr: []string{},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
		},
	}, // End of all tests.

	/*
	 * Test for the function PrintBlank().
	 */
	{
		testingFunction: func(in TestList) string {
			world := NewWorld()
			out := world.PrintBlank()
			return out
		},
		tests: []TestList{
			{
				testName: "map_pretty-print_blank_00",
				isMulti:  false,
				inputArr: []string{},
				expectedValue: // NOTE: this comment breaks the line.
				"┌────────────────────────┐" + NL +
					"│1-987654321 123456789+12│" + NL +
					"├────────────────────────┤" + NL +
					"│    _,--._  _._.--.--.._│" + NL +
					"│=.--'=_',-,:`;_      .,'│" + NL +
					"│,-.  _.)  (``-;_   .'   │" + NL +
					"│   '-:_    `) ) .''=.   │" + NL +
					"│     ) )    ()'    ='   │" + NL +
					"│     |/            (_) =│" + NL +
					"│     -                  │" + NL +
					"├────────────────────────┤" + NL +
					"│                        │" + NL +
					"└────────────────────────┘",
			},
		},
	}, // End of all tests.

	/*
	 * Test for the function PrintDemo().
	 */
	{
		testingFunction: func(in TestList) string {
			world := NewWorld()
			out := world.PrintDemo()
			return out
		},
		tests: []TestList{
			{
				testName: "map_pretty-print_demo_00",
				isMulti:  true,
				inputArr: []string{},
				expectedValue: // NOTE: this comment breaks the line.
				"┌────────────────────────┐" + NL +
					"│1-987654321 123456789+12│" + NL +
					"├───────────▼────────────┤" + NL +
					"│    _,--._  _._.--.--.._│" + NL +
					"▶=.--'=_',-,▣`;_      .,'◀" + NL +
					"│,-.  _.)  (``-;_   .'   │" + NL +
					"│   '-:_    `) ) .''=.   │" + NL +
					"│     ) )    ()'    ='   │" + NL +
					"│     |/            (_) =│" + NL +
					"│     -                  │" + NL +
					"├───────────▲────────────┤" + NL +
					"│   ☼            ●       │" + NL +
					"└────────────────────────┘",
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
