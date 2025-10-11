package dataprov

import (
	"fmt"
	"testing"
	"time"

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
	// Definition of test data.
	testData := []struct {
		input string
		exp   string
	}{
		{
			input: "2025-10-03 22:12:16",
			exp:   "moon lat: -001.8" + NL + "moon lon: +147.8",
			// data source: https://doncarona.tamu.edu/cgi-bin/moon?current=0&jd=
			// lat:   -1.807 (south).
			// lon: +147.767 (west).
		},
		/* TODO: improve calculation of moon's geocentric lat+lon values (ie sublunar position).
		{
			input: "2025-10-10 21:33:42",
			exp:   "moon lat: +005.1" + NL + "moon lon: +070.0",
			// data source: https://doncarona.tamu.edu/cgi-bin/moon?current=0&jd=
			// lat:  +5.111 (north).
			// lon: +69.992 (west).
		},
		*/
	}
	// Providers under test.
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
		for _, data := range testData {
			t.Run(test.name, func(t *testing.T) {
				err0 := test.provider.SetTime(data.input)
				if err0 != nil {
					t.Fatalf("Setup failed: %v", err0)
				}
				lat, lon := test.provider.MoonsGeocentricCoords()
				got := fmt.Sprintf(
					"moon lat: %+06.01f"+NL+
						"moon lon: %+06.01f",
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

func TestTableDrivenOfSunsetDataProviders(t *testing.T) {
	// Definition of test data.
	testTolerance := 6 * time.Minute // TODO: decrease tolerance to 1 minute.
	testData := []struct {
		inputTime string
		// NOTE: maybe use unit.Angle for lat and lon instead of long64's.
		inputLat   float64
		inputLon   float64
		expSunrise string
		expSunset  string
		expPrint   string
	}{
		// TODO: make sure not to use elevation of observation for sunrise/sunset in data source, but calculate for some sort of zero-level (eg MSL, mean sea level).
		// TODO: make sure to also show the year-month-day for each time (to not confuse it with a time of the next or previous day).
		{
			inputTime: "2025-10-14 20:40:50",
			// somewhere close to or in texas.
			inputLat:   +32.66578, // lat: N 32°39'56.8'' or +32.66578°.
			inputLon:   -96.28518, // lon: W 96°17'6.65'' or -96.28518°.
			expSunrise: "2025-10-14 12:26:11",
			expSunset:  "2025-10-14 23:56:50",
			expPrint:   "sunrise: 12:26:11" + NL + "sunset:  23:56:50", // UTC.
			// exp:   "sunrise: 07:26:11" + NL + "sunset:  18:56:50", // local time.
			// data source: https://www.suncalc.org/#/32.6658,-96.2852,3/2025.10.11/00:41/1/3
		},
		{
			inputTime: "2025-10-14 20:40:50",
			// houston, texas.
			inputLat:   +29.75000, // lat: N 29°45'00.0'' or +29.75000°.
			inputLon:   -95.35000, // lon: W 95°21'0.00'' or -95.35000°.
			expSunrise: "2025-10-14 12:21:00",
			expSunset:  "2025-10-14 23:55:00",
			expPrint:   "sunrise: 12:21:00" + NL + "sunset:  23:55:00", // UTC.
			// exp:   "sunrise: 06:21:11" + NL + "sunset:  17:55:00", // local in houston, tx.
			// data source: https://gml.noaa.gov/grad/solcalc/sunrise.html
		},
	}
	// Providers under test.
	tests := []struct {
		name     string
		provider SunsetSunriseDataProviderInterface
	}{
		{
			name:     "data-provider_sunset_globe_00",
			provider: &OsmanSunsetSunriseDataProvider{},
		},
	}
	// Other stuff.
	simpleTimeLayout := "2006-01-02 15:04:05"
	timeFormat := "15:04:05"
	// Loop over test cases.
	for _, test := range tests {
		for _, data := range testData {
			t.Run(test.name, func(t *testing.T) {
				err0 := test.provider.SetTime(data.inputTime)
				if err0 != nil {
					t.Fatalf("Setup failed: %v", err0)
				}
				sunrise := test.provider.Sunrise(data.inputLat, data.inputLon)
				sunset := test.provider.Sunset(data.inputLat, data.inputLon)
				expSr, _ := time.ParseInLocation(simpleTimeLayout, data.expSunrise, time.UTC)
				expSs, _ := time.ParseInLocation(simpleTimeLayout, data.expSunset, time.UTC)
				failSr := !isWithinInterval(testTolerance, sunrise, expSr)
				failSs := !isWithinInterval(testTolerance, sunset, expSs)
				if failSr || failSs {
					got := fmt.Sprintf("sunrise: %s"+NL+"sunset:  %s", sunrise.Format(timeFormat), sunset.Format(timeFormat))
					t.Errorf("In '%s':\n", test.name)
					diff := godiff.CDiff(data.expPrint, got)
					t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", data.expPrint, got)
					t.Errorf("exp/got:\n%s\n", diff)
					t.Errorf("diff sunrise: %s (HH:MM:SS)", prettyPrintTimeDiff(sunrise, expSr))
					t.Errorf("diff sunset:  %s (HH:MM:SS)", prettyPrintTimeDiff(sunset, expSs))
				}
			})
		}
	}
}

func isWithinInterval(interval time.Duration, t1, t2 time.Time) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff <= interval
}

func prettyPrintTimeDiff(t1 time.Time, t2 time.Time) string {
	timeDiff := t1.Sub(t2)
	totalSeconds := int(timeDiff.Seconds())
	h := totalSeconds / 3600
	m := (totalSeconds % 3600) / 60
	s := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
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
		{
			input: "2025-10-10 21:24:41",
			exp:   "sun lat: -007.0" + NL + "sun lon: +144.5",
			// data source: https://www.timeanddate.com/
			// lat:   -7.000 degrees (south).
			// lon: +144.466 degrees (west).
		},
		{
			input: "2001-01-01 13:01:01",
			exp:   "sun lat: -023.0" + NL + "sun lon: +014.3",
			// data source: https://www.timeanddate.com/worldclock/sunearth.html?day=01&month=01&year=2001&hour=13&min=01&sec=01&n=&ntxt=&earth=0
			// lat: -22.966 degrees (south).
			// lon: +14.333 degrees (west).
		},
		{
			input: "2100-01-01 13:01:01",
			exp:   "sun lat: -023.0" + NL + "sun lon: +014.4",
			// data source: https://www.timeanddate.com/worldclock/sunearth.html?day=1&month=1&year=2100&hour=13&min=1&sec=1&n=&ntxt=&earth=0
			// lat: -22.966 degrees (south).
			// lon: +14.400 degrees (west).
		},
	}
	// Providers under test.
	tests := []struct {
		name     string
		provider SunDataProviderInterface
	}{
		{
			name:     "data-provider_sun_globe_00",
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
