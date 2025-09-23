package geomain

import (

	// this is a test.
	"testing"

	// printing and formatting.
	"fmt"

	// other imports.
	"github.com/kraasch/godiff/godiff"
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

var suites = []TestSuite{
	/*
	* Test for the function PrintEarth().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			geodata := New()
			out = geodata.PrintData()
			return out
		},
		tests: []TestList{
			{
				testName: "earth_pretty-print_overview_00",
				isMulti:  true,
				inputArr: []string{},
				expectedValue: // NOTE: this comment breaks the line.
				"┌────────────────────────┐" + NL +
					"│1-987654321 123456789+12│" + NL +
					"├────────────▼───────────┤" + NL +
					"│    _,--._  _._.--.--.._│" + NL +
					"│=.--'=_',-,:`;_      .,'│" + NL +
					"│,-.  _.)  (``-;_   .'   │" + NL +
					"▶   '-:_    `▣ ) .''=.   ◀" + NL +
					"│     ) )    ()'    ='   │" + NL +
					"│     |/            (_) =│" + NL +
					"│     -                  │" + NL +
					"├────────────▲───────────┤" + NL +
					"│            ☼           │" + NL +
					"└────────────────────────┘" + NL +
					" ● age:       19.76 days" + NL +
					" ● phase:     Waning Gibbous (◐)" + NL +
					" ● dist.:     371578 km" + NL +
					" ● illum.:    74%" + NL +
					" next new ●:  10.6 days (0001-01-11, Thu)" + NL +
					" next full ●: 25.2 days (0001-01-26, Fri)" + NL +
					" ☼ rise:      01:10 h" + NL +
					" ☼ set:       13:17 h",
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
