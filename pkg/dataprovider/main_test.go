package dataprov

import (
	"fmt"
	"testing"

	// other imports.
	godiff "github.com/kraasch/godiff/godiff"
)

var NL = fmt.Sprintln()

func TestTableDrivenOfGeneralDataProviders(t *testing.T) {
	input := "2025-10-03 22:12:16"
	exp := "julian date: 2460952.425" // expected output string.
	// data source: https://www.calendarlabs.com/julian-date-converter
	tests := []struct {
		name     string
		provider GeneralDataProviderInterface
	}{
		{
			name:     "data-provider_moon_keys_00",
			provider: &GlobeGeneralDataProvider{},
		},
		{
			name:     "data-provider_moon_sampa_00",
			provider: &KeysGeneralDataProvider{},
		},
	}
	// Loop over test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err0 := test.provider.SetTime(input)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			jd := test.provider.JulianDate()
			got := fmt.Sprintf("julian date: %11.3f", jd)
			if exp != got {
				t.Errorf("In '%s':\n", test.name)
				diff := godiff.CDiff(exp, got)
				t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
				t.Errorf("exp/got:\n%s\n", diff)
			}
		})
	}
}

func TestTableDrivenOfMoonDataProviders(t *testing.T) {
	input := "2025-10-03 22:12:16"
	// data source: https://doncarona.tamu.edu/cgi-bin/moon?current=0&jd=
	exp := // expected output string.
	// Geocentric Latitude  -1.807
	// Geocentric Longitude  147.767
	"moon lat: -001.8" + NL +
		"moon lon: +147.8"
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
	// Loop over test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err0 := test.provider.SetTime(input)
			if err0 != nil {
				t.Fatalf("Setup failed: %v", err0)
			}
			lat, lon := test.provider.MoonsGeocentricCoords()
			got := fmt.Sprintf(
				"moon lat: %+06.01f"+NL+
					"moon lon: %+06.01f",
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

func TestTableDrivenOfSunDataProviders(t *testing.T) {
	// Definition of test data.
	testData := []struct {
		input string
		exp   string
	}{
		{
			input: "2025-10-14 20:40:50",
			exp:   "sun lat: -008.5" + NL + "sun lon: +133.7",
			// data source: https://www.timeanddate.com/worldclock/sunearth.html?day=14&month=10&year=2025&hour=20&min=40&sec=50&n=&ntxt=&earth=0
			// lat: -8.483 degrees (south).
			// lon: 133.733 degrees (west).
		},
		{
			input: "2010-10-20 20:30:49",
			exp:   "sun lat: -010.5" + NL + "sun lon: +131.5",
			// data source: https://www.timeanddate.com/
			// lat: -10.516 degrees (south).
			// lon: +131.533 degrees (west).
		},
		{
			input: "2000-07-07 08:30:49",
			exp:   "sun lat: +022.5" + NL + "sun lon: -053.5",
			// data source: https://www.timeanddate.com/
			// lat: +22.533 (north).
			// lon: -53.533 (east).
		},
	}
	// Providers under test.
	tests := []struct {
		name     string
		provider SunDataProviderInterface
	}{
		{
			name:     "data-provider_sun_keys_00",
			provider: &GlobeSunDataProvider{},
		},
	}
	// Loop over test cases.
	for _, test := range tests {
		for _, data := range testData {
			t.Run(test.name, func(t *testing.T) {
				err0 := test.provider.SetTime(data.input)
				if err0 != nil {
					t.Fatalf("Setup failed: %v", err0)
				}
				lat, lon := test.provider.SunsGeocentricCoords()
				got := fmt.Sprintf(
					"sun lat: %+06.01f"+NL+
						"sun lon: %+06.01f",
					lat, lon)
				if data.exp != got {
					t.Errorf("In '%s':\n", test.name)
					diff := godiff.CDiff(data.exp, got)
					t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", data.exp, got)
					t.Errorf("exp/got:\n%s\n", diff)
				}
			})
		}
	}
}
