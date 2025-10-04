package astro

import (
	"fmt"
	"testing"
	"time"

	// other imports.

	godiff "github.com/kraasch/godiff/godiff"
)

var (
	NL               = fmt.Sprintln()
	simpleTimeLayout = "2006-01-02 15:04:05"
)

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
	 * julian date = julian day + fraction of the day.
	 *  NOTE: julian day + 0.25 is plus 6 hours past noon, not past midnight!
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			theTime, err0 := time.ParseInLocation(simpleTimeLayout, inputTime, time.UTC)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			jd := JulianDate(theTime)
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

	/*
	 * Test JulianEphemerisDay().
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			theTime, err0 := time.ParseInLocation(simpleTimeLayout, inputTime, time.UTC)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			jd := JulianDate(theTime)
			jde := JulianEphemerisDay(jd, theTime)
			return fmt.Sprintf("julian ephemeris day: %11.3f", jde)
		},
		tests: []TestList{
			{
				testName:      "astro_basic-calculations_time_00",
				isMulti:       true,
				inputArr:      []string{"2025-10-03 22:12:16"},     // input.
				expectedValue: "julian ephemeris day: 2460952.426", // output.
				// TODO: find external data source for this data.
			},
		},
	},

	/*
		* Test moon position.
		{
			testingFunction: func(in TestList, t *testing.T) string {
				inputTime := in.inputArr[0]
				theTime, err0 := time.ParseInLocation(simpleTimeLayout, inputTime, time.UTC)
				if err0 != nil {
					t.Fatalf("Setup failed: %v", err0)
				}
				lat, lon := MoonPosition(theTime)
				return fmt.Sprintf("lat: %+06.01f, lon: %+06.01f", lat, lon)
			},
			tests: []TestList{
				{
					// Test data source: https://doncarona.tamu.edu/cgi-bin/moon?current=0&jd=
					// Time 2025-10-03T22:12:15.895 UTC
					// Geocentric Latitude  -1.807
					// Geocentric Longitude  327.767
					testName:      "astro_more-calculations_moon_00",
					isMulti:       true,
					inputArr:      []string{"2025-10-03 22:12:16"}, // input time.
					expectedValue: "lat: -001.8, lon: +147.8",      // output coordinates.
				},
			},
		},
	*/
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
