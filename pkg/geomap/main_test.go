package geomap

import (

	// this is a test.
	"strconv"
	"testing"

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
	 * Test for the function SubLine().
	 */
	{
		testingFunction: func(in TestList) string {
			lon0, err0 := strconv.ParseFloat(in.inputArr[0], 64)
			lon1, err1 := strconv.ParseFloat(in.inputArr[1], 64)
			if err0 != nil {
				return "error in type converstion within the test: first float."
			}
			if err1 != nil {
				return "error in type converstion within the test: second float."
			}
			line := NewSubLine()
			reserr0 := line.AddMoon(lon0)
			reserr1 := line.AddSun(lon1)
			out := line.Line
			if reserr0 != nil || reserr1 != nil {
				return "error getting a result."
			}
			return out
		},
		tests: []TestList{
			{
				testName: "map_line-under-display_00",
				isMulti:  true,
				inputArr: []string{
					"-180.0", // longitude, ie (").
					"+180.0", // longitude, ie (").
				},
				expectedValue:// NOTE: this comment breaks the line.
				//"1-987654321 123456789+12"
				"●                      ☼",
			},
		},
	},

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
				isMulti:  true,
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
					"▣  '-:_    `) ) .''=.   " + NL + // lat. +20 to -20. = 40 degrees.
					"     ) )    ()'    ='   " + NL + // lat. -20 to -35. = 15 degrees.
					"     |/            (_) =" + NL + // lat. -35 to -50. = 15 degrees.
					"     -                  ", //       lat. -50 to -90. = 40 degrees.
				// A          A           A
				// |          |           |
				// |        0 to 15       | ==> 360/24 = 15        // 24 columns, 360 degrees on the globe.
				// |     -7.5 to 7.5      |                        // thus: 0-15 degrees are the first column, and so on.
				// |                      | func lon2col( lon ):   // write a function to translate latitude to columns.
				// ._ -180(E)    +165(W) _.   lon = lon+180        // start with 0 at the left.
				//  to-165(E)  to+180(W)      return int(lon / 15) // divide without rest.

				// NOTE: idea: center around each of the 24 zones (with each 15 degrees) by subtracting half of 15 = 7.5.
				// NOTE: implement another time.
				// column |       +180 |          -7.5 |
				// -------+------------+---------------+
				// 00     |   0 to  15 |  -7.5 to   7.5|
				// 01     |  15 to  30 |   7.5 to  22.5|
				// 02     |  30 to  45 |  22.5 to  37.5|
				// 03     |  45 to  60 |  37.5 to  52.5|
				// 04     |  60 to  75 |  52.5 to  67.5|
				// 05     |  75 to  90 |  67.5 to  82.5|
				// 06     |  90 to 105 |  82.5 to  97.5|
				// 07     | 105 to 120 |  97.5 to 112.5|
				// 08     | 120 to 135 | 112.5 to 127.5|
				// 09     | 135 to 150 | 127.5 to 142.5|
				// 10     | 150 to 165 | 142.5 to 157.5|
				// 11     | 165 to 180 | 157.5 to 172.5|
				// 12     | 180 to 195 | 172.5 to 187.5|
				// 13     | 195 to 210 | 187.5 to 202.5|
				// 14     | 210 to 225 | 202.5 to 217.5|
				// 15     | 225 to 240 | 217.5 to 232.5|
				// 16     | 240 to 255 | 232.5 to 247.5|
				// 17     | 255 to 270 | 247.5 to 262.5|
				// 18     | 270 to 285 | 262.5 to 277.5|
				// 19     | 285 to 300 | 277.5 to 292.5|
				// 20     | 300 to 315 | 292.5 to 307.5|
				// 21     | 315 to 330 | 307.5 to 322.5|
				// 22     | 330 to 345 | 322.5 to 337.5|
				// 23     | 345 to 360 | 337.5 to 352.5|
			},
			{
				testName: "map_pretty-print_coord_01",
				isMulti:  true,
				inputArr: []string{
					"0.0",    // latitude, ie (=).
					"+179.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.  ▣" + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_02",
				isMulti:  true,
				inputArr: []string{
					"90.0",   // latitude, ie (=).
					"+179.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--..▣" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_03",
				isMulti:  true,
				inputArr: []string{
					"-90.0",  // latitude, ie (=).
					"+179.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                 ▣",
			},
			{
				testName: "map_pretty-print_coord_allow-full180_00",
				isMulti:  true,
				inputArr: []string{
					"0.0",    // latitude, ie (=).
					"+180.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.  ▣" + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_04",
				isMulti:  true,
				inputArr: []string{
					"0.0", // latitude, ie (=).
					"0.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `▣ ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_05",
				isMulti:  true,
				inputArr: []string{
					"90.0", // latitude, ie (=).
					"0.0",  // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  ▣._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_06",
				isMulti:  true,
				inputArr: []string{
					"-90.0", // latitude, ie (=).
					"0.0",   // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -      ▣           ",
			},
			{
				testName: "map_pretty-print_coord_07",
				isMulti:  true,
				inputArr: []string{
					"90.0",   // latitude, ie (=).
					"-180.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"▣   _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"     -                  ",
			},
			{
				testName: "map_pretty-print_coord_08",
				isMulti:  true,
				inputArr: []string{
					"-90.0",  // latitude, ie (=).
					"-180.0", // longitude, ie (").
				},
				expectedValue: // NOTE: this comment breaks the line.
				"    _,--._  _._.--.--.._" + NL +
					"=.--'=_',-,:`;_      .,'" + NL +
					",-.  _.)  (``-;_   .'   " + NL +
					"   '-:_    `) ) .''=.   " + NL +
					"     ) )    ()'    ='   " + NL +
					"     |/            (_) =" + NL +
					"▣    -                  ",
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
