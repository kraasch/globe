package geoshow

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	// other imports.
	"github.com/kraasch/godiff/godiff"
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

var suites = []TestSuite{
	/*
	 * Test for the function SunRiseAndSet().
	 */
	{
		testingFunction: func(in TestList) string {
			lat, err0 := strconv.ParseFloat(in.inputArr[0], 64)
			lon, err1 := strconv.ParseFloat(in.inputArr[1], 64)
			time, err2 := time.Parse(dateLayout, in.inputArr[2])
			if err0 != nil {
				return "error in type converstion within the test: first float."
			}
			if err1 != nil {
				return "error in type converstion within the test: second float."
			}
			if err2 != nil {
				return "error in type converstion within the test: date."
			}
			out := SunRiseAndSet(lat, lon, time)
			return out
		},
		tests: []TestList{
			{
				testName: "sun_sunrise+sunset_calculate_00",
				isMulti:  false,
				inputArr: []string{
					"43.65", "-79.38", // Toronto, Canada.
					"2000-01-01", // Some date.
				},
				// expectedValue: "sunrise: 12:51, sunset: 21:51", // NOTE: as UTC.
				expectedValue: "sunrise: 07:51, sunset: 16:51", // NOTE: as GMT-5 (in Toronto, Canada).
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
