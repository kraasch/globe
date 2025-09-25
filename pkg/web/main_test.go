package web

import (
	"fmt"
	"testing"

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

var suites = []TestSuite{ // All tests.

	/*
	 * Test for the function Toast().
	 */
	{
		testingFunction: func(in TestList) string {
			out := Toast()
			return out
		},
		tests: []TestList{
			{
				testName: "auto_basic_test_00",
				isMulti:  true,
				inputArr: []string{
					"2000-01-01", // Some date.
				},
				// expectedValue: "zone: UTC, lon: 53.48, lat: 10.22",
				expectedValue: "Toast!",
				// TODO: note that this test's outcome depends on where the executing machine is located.
				// TODO: remove or rewrite this test to make it independent from the machines location.
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
