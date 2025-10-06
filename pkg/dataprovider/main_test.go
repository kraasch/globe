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

// ////////////////////////////////////////////
// TEST SUITE WITH MULIT-LINE STRING TESTS. //
// ////////////////////////////////////////////
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

// /////////////////////
// TABLE-DRIVEN TESTS //
// /////////////////////
func TestTableDrivenOfMoonDataProviders(t *testing.T) {
	input := "2025-10-03 22:12:16"
	exp := // expected output string.
	"lat: -001.8" + NL +
		"lon: +147.8"
	tests := []struct {
		name     string
		provider MoonDataProviderInterface
	}{
		{
			name:     "data-provider_moon_keys_00",
			provider: &KeysMoonDataProvider{},
		},
		{
			name:     "data-provider_moon_sampa_00",
			provider: &SampaMoonDataProvider{},
		},
	}
	// Loop over test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err0 error = nil
			switch p := test.provider.(type) {
			case *KeysMoonDataProvider:
				err0 = p.SetTime(input)
			case *SampaMoonDataProvider:
				err0 = p.SetTime(input)
			}
			// provider := dataprov.KeysMoonDataProvider{}
			// _ = provider.SetTime(t.Format("2006-01-02 15:04:05"))
			// return provider.GeocentricCoords()
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			lat, lon := test.provider.GeocentricCoords()
			got := fmt.Sprintf(
				"lat: %+06.01f"+NL+
					"lon: %+06.01f",
				lat, lon)
			if exp != got {
				t.Errorf("In '%s':\n", test.name)
				diff := godiff.CDiff(exp, got)
				t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
				t.Errorf("exp/got:\n%s\n", diff)
			}
		})
	}
}
