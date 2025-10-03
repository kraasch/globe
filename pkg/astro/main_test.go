package astro

import (
	"fmt"
	"testing"
	"time"

	// other imports.

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
	testingFunction func(in TestList, t *testing.T) string
	tests           []TestList
}

var suites = []TestSuite{ // All tests.

	/*
	 * Test JulianDate().
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			simpleTimeLayout := "2006-01-02 15:04:05"
			theTime, err0 := time.Parse(simpleTimeLayout, inputTime)
			jd := JulianDate(theTime)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			return fmt.Sprintf("julian date: %11.3f", jd)
		},
		tests: []TestList{
			{
				testName:      "astro_basic-calculations_time_00",
				isMulti:       true,
				inputArr:      []string{"2025-10-03 22:12:16"}, // input.
				expectedValue: "julian date: 2460952.425",      // output.
				// data source: https://www.calendarlabs.com/julian-date-converter
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
				got := suite.testingFunction(test, t)
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
