package geoshow

import (
	"fmt"
	"time"

	mp "github.com/janczer/goMoonPhase"
	sr "github.com/nathan-osman/go-sunrise"
)

func MoonPhase(date time.Time) string {
	phase := mp.New(date)
	return fmt.Sprintf("%#v", phase)
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
	return fmt.Sprintf("sunrise: %02d:%02d, sunset: %02d:%02d", localRise.Hour(), localRise.Minute(), localSet.Hour(), localSet.Minute())
}
