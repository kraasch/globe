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
	tests := []struct {
		name     string
		provider GeneralDataProviderInterface
	}{
		{
			name:     "data-provider_moon_keys_00",
			provider: &GeoGeneralDataProvider{},
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
	// Loop over test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err0 := test.provider.SetTime(input)
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
