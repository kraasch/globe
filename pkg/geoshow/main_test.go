package geoshow

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
			out = Toast()
			return out
		},
		tests: []TestList{
			{
				testName: "earth_pretty-print_overview_00",
				isMulti:  false,
				inputArr: []string{},
				expectedValue:// NOTE: this comment breaks the line.
				"Toast!",
				//"┌────────────────────────┐" + NL +
				// "│1 9876-4321 1234+6789 12│" + NL +
				// "├───────────▼────────────┤" + NL +
				// "│    _,--._  _._.--.--.._│" + NL +
				// "▶=.--'=_',-,▣`;_      .,'◀" + NL +
				// "│,-.  _.)  (``-;_   .'   │" + NL +
				// "│   '-:_    `) ) v''=.   │" + NL +
				// "│     ) )    ()'    ='   │" + NL +
				// "│     |/            (_) =│" + NL +
				// "├───────────▲────────────┤" + NL +
				// "│   ☼            ●       │" + NL +
				// "└────────────────────────┘",
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
