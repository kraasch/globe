// Package geocalc does the main under-the-hood calculations.
package geocalc

import (
	"fmt"
	"time"

	mp "github.com/janczer/goMoonPhase"
	sr "github.com/nathan-osman/go-sunrise"
)

const (
	DAYFORMAT       = "2006-01-02, Mon" // A Go-style format string for dates.
	DAYFORMATMIDDLE = "2006-01-02"      // A Go-style format string for dates.
	DAYFORMATSHORT  = "06-01-02"        // A Go-style format string for dates.
	TIMEFORMAT      = "15:04"           // A Go-style format string for time (with leading zero: ie. 07:45h )
)

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
	// TODO: convert lat and lon to timezone string.
	localization := time.FixedZone("GMT-5", -5*60*60) // TODO: extract this into a universal function which can return any time zone depending on lon and lat.
	localRise := rise.In(localization)
	localSet := set.In(localization)
	// round up for more than 30 seconds.
	if localRise.Second() > 30 {
		localRise = localRise.Add(time.Minute)
	}
	if localSet.Second() > 30 {
		localSet = localSet.Add(time.Minute)
	}
	return fmt.Sprintf(" ☼ rise:    %02d:%02d h\n ☼ set:     %02d:%02d h", localRise.Hour(), localRise.Minute(), localSet.Hour(), localSet.Minute())
}
