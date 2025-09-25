package geotop

import (

	// this is a test.
	"testing"

	// printing and formatting.
	"fmt"

	// other imports.
	util "github.com/kraasch/geo/pkg/testutil"
	godiff "github.com/kraasch/godiff/godiff"
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
	* Test for the function PrintDataVertically().
	* Also kind of tests PrintDataHorizontally(), because it uses the same data.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			geodata := New()
			out = geodata.PrintDataVertically()
			out = util.Anonymize(out)
			return out
		},
		tests: []TestList{
			{
				testName: "earth_pretty-print_overview_00",
				isMulti:  true,
				inputArr: []string{},
				expectedValue: // NOTE: this comment breaks the line.
				"┌────────────────────────┐" + NL +
					util.AnonymizeNumberStrings("│1-987654321 123456789+12│") + NL +
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
					" ○ phase:   Waning Gibbous" + NL +
					util.Anonymize(" ○ age:     19.76 days (◐)") + NL +
					util.Anonymize(" ○ dist.:   371578 km") + NL +
					util.Anonymize(" ○ illum.:  74%") + NL +
					util.Anonymize(" ○ new in:  05.2d 2000-01-06") + NL +
					util.Anonymize(" ○ full in: 20.2d 2000-01-06") + NL +
					util.Anonymize(" □ time:    22:02 h") + NL +
					util.Anonymize(" 🜃 utc:     20:02 h") + NL +
					util.Anonymize(" ☼ rise:    01:10 h") + NL +
					util.Anonymize(" ☼ set:     13:17 h"),
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
