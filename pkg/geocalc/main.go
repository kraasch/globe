// Package geocalc does the main under-the-hood calculations.
package geocalc

import (
	"fmt"
	"time"

	mp "github.com/janczer/goMoonPhase"
	gw "github.com/kraasch/geo/pkg/geoweb"
	sr "github.com/nathan-osman/go-sunrise"
	tzf "github.com/ringsaturn/tzf"
)

const (
	DAYFORMAT       = "2006-01-02, Mon" // A Go-style format string for dates.
	DAYFORMATMIDDLE = "2006-01-02"      // A Go-style format string for dates.
	DAYFORMATSHORT  = "06-01-02"        // A Go-style format string for dates.
	TIMEFORMAT      = "15:04"           // A Go-style format string for time (with leading zero: ie. 07:45h )
)

var tzFinder tzf.F

var webBuf gw.WebBuffer = gw.NewWebBuffer()

func initTzf() {
	if tzFinder == nil {
		var err error
		tzFinder, err = tzf.NewDefaultFinder()
		if err != nil {
			panic(err)
		}
	}
}

func timeToUtcOffset(t time.Time) string {
	// Get the zone name and offset in seconds
	zoneName, offsetSeconds := t.Zone()
	// Calculate hours and minutes from seconds
	sign := "+"
	if offsetSeconds < 0 {
		sign = "-"
		offsetSeconds = -offsetSeconds
	}
	hours := offsetSeconds / 3600
	minutes := (offsetSeconds % 3600) / 60
	// Format as UTC+X or UTC-X
	offsetStr := fmt.Sprintf("UTC%s%d", sign, hours)
	if minutes != 0 {
		offsetStr += fmt.Sprintf(":%02d", minutes)
	}
	return fmt.Sprintf("%s (%s)", offsetStr, zoneName)
}

func ConvertLatAndLonToTimezone(lat, lon float64) string {
	initTzf()                                // init tzFinder variable.
	tz := tzFinder.GetTimezoneName(lon, lat) // NOTE: Takes longitude-latitude order.
	return tz
}

func LatAndLonAndTz() string {
	lat, lon := webBuf.GetCoords()
	tz := ConvertLatAndLonToTimezone(lat, lon)
	now := time.Now()
	utcOffset := timeToUtcOffset(now)
	return fmt.Sprintf(" □ lat+lon: %.2f, %.2f\n □ zone:    %s\n □ offset:  %s\n", lat, lon, tz, utcOffset)
}

func LatAndLon() string {
	lat, lon := webBuf.GetCoords()
	return fmt.Sprintf(" ▣ lat+lon: %.2f, %.2f", lat, lon)
}

func LocalAndUtcTime() string {
	now := time.Now()
	formattedNow := now.Format(TIMEFORMAT)
	utc := now.UTC()
	formattedUtc := utc.Format(TIMEFORMAT)
	return fmt.Sprintf(" □ time:    %s h\n 🜃 utc:     %s h", formattedNow, formattedUtc)
}

func PhaseToText(phase float64) string {
	stage := ""
	switch {
	case phase >= 0 && phase < 0.125:
		stage = "New Moon"
	case phase >= 0.125 && phase < 0.25:
		stage = "Waxing Crescent"
	case phase >= 0.25 && phase < 0.375:
		stage = "First Quarter"
	case phase >= 0.375 && phase < 0.5:
		stage = "Waxing Gibbous"
	case phase >= 0.5 && phase < 0.625:
		stage = "Full Moon"
	case phase >= 0.625 && phase < 0.75:
		stage = "Waning Gibbous"
	case phase >= 0.75 && phase < 0.875:
		stage = "Last Quarter"
	case phase >= 0.875 && phase < 1:
		stage = "Waning Crescent"
	}
	return stage
}

func PhaseToSymbol(phase float64) string {
	stage := ""
	switch {
	case phase >= 0 && phase < 0.125:
		stage = "○"
	case phase >= 0.125 && phase < 0.25:
		stage = "❩"
	case phase >= 0.25 && phase < 0.375:
		stage = "◗"
	case phase >= 0.375 && phase < 0.5:
		stage = "◑"
	case phase >= 0.5 && phase < 0.625:
		stage = "●"
	case phase >= 0.625 && phase < 0.75:
		stage = "◐"
	case phase >= 0.75 && phase < 0.875:
		stage = "◖"
	case phase >= 0.875 && phase < 1:
		stage = "❨"
	}
	return stage
}

func unixToDate(unixDate float64) time.Time {
	seconds := int64(unixDate)
	nanoseconds := int64((unixDate - float64(seconds)) * 1e9)
	t := time.Unix(seconds, nanoseconds)
	return t
}

func diffInDays(date1, date2 time.Time) float64 {
	diff := date2.Sub(date1)
	return diff.Hours() / 24
}

func MoonPhaseVerbose(date time.Time) string {
	phase := mp.New(date)
	nextNew := unixToDate(phase.NextNewMoon())   // in unix format (ms since 1970) -- i guess.
	nextFull := unixToDate(phase.NextFullMoon()) // in unix format (ms since 1970) -- i guess.
	return fmt.Sprintf(
		"The moon is %.2f days old, and is therefore in %s phase (%s).\nIt is %.0f km from the centre of the Earth.\nIt is %.0f%% illuminated.\nThe next new moon is in %.1f days (%v).\nThe next full moon is in %.1f days (%v).",
		phase.Age(),                  // age in days -- i guess.
		PhaseToText(phase.Phase()),   // convert moonphase (0-1 value) to text -- i guess.
		PhaseToSymbol(phase.Phase()), // convert moonphase (0-1 value) to symbol -- i guess.
		phase.Distance(),             // distance from earth in km -- i guess.
		phase.Illumination()*100,     // illumination between 0 and 1 -- i guess.
		diffInDays(date, nextNew),
		nextNew.Format(DAYFORMAT),
		diffInDays(date, nextFull),
		nextFull.Format(DAYFORMAT),
	)
}

func MoonPhase(date time.Time) string {
	phase := mp.New(date)
	nextNew := unixToDate(phase.NextNewMoon())   // in unix format (ms since 1970) -- i guess.
	nextFull := unixToDate(phase.NextFullMoon()) // in unix format (ms since 1970) -- i guess.
	return fmt.Sprintf(
		` ○ phase:   %s
 ○ age:     %.2f days (%s)
 ○ dist.:   %.0f km
 ○ illum.:  %.0f%%
 ○ new in:  %04.1fd %s
 ○ full in: %04.1fd %s`,
		PhaseToText(phase.Phase()),   // convert moonphase (0-1 value) to text -- i guess.
		phase.Age(),                  // age in days -- i guess.
		PhaseToSymbol(phase.Phase()), // convert moonphase (0-1 value) to symbol -- i guess.
		phase.Distance(),             // distance from earth in km -- i guess.
		phase.Illumination()*100,     // illumination between 0 and 1 -- i guess.
		diffInDays(date, nextNew),
		nextNew.Format(DAYFORMATMIDDLE),
		diffInDays(date, nextFull),
		nextFull.Format(DAYFORMATMIDDLE),
	)
}

func SunRiseAndSet(lat, lon float64, date time.Time) string {
	rise, set := sr.SunriseSunset(lat, lon, date.Year(), date.Month(), date.Day())
	return fmt.Sprintf(" ☼ rise:    %02d:%02d h\n ☼ set:     %02d:%02d h", rise.Hour(), rise.Minute(), set.Hour(), set.Minute())
}
