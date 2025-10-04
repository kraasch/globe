package dataprov

import (
	"fmt"
	"testing"

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
	* Test general data providers: geo.
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			provider := NewGeoGeneralDataProvider()
			err0 := provider.SetTime(inputTime)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			jd := provider.JulianDate()
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
	* Test general data providers: keys.
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			provider := NewKeysGeneralDataProvider()
			err0 := provider.SetTime(inputTime)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			jd := provider.JulianDate()
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
	* Test moon data providers: sampa.
	 */
	{
		testingFunction: func(in TestList, t *testing.T) string {
			inputTime := in.inputArr[0]
			provider := NewSampaMoonDataProvider()
			err0 := provider.SetTime(inputTime)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			lat, lon := provider.GeocentricCoords()
			return fmt.Sprintf("lat: %+06.01f, lon: %+06.01f", lat, lon)
		},
		tests: []TestList{
			{
				// Test data source: https://doncarona.tamu.edu/cgi-bin/moon?current=0&jd=
				// Time 2025-10-03T22:12:15.895 UTC
				// Geocentric Latitude  -1.807
				// Geocentric Longitude  327.767
				testName:      "data-provider_moon_basic_00",
				isMulti:       true,
				inputArr:      []string{"2025-10-03 22:12:16"}, // input time.
				expectedValue: "lat: -001.8, lon: +147.8",      // output coordinates.
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
