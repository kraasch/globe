// Package astro is rolling geo's own calculations.
package astro

import (
	"time"
)

func JulianDate(t time.Time) float64 {
	// JulianDate should only take a UTC time. // TODO: fix.
	// t = t.UTC()
	year := t.Year()
	month := int(t.Month())
	day := float64(t.Day())
	// Calculate the Julian Day Number.
	if month <= 2 { // Adjust months and years for January and February.
		year -= 1
		month += 12
	}
	A := int(year / 100)
	B := 2 - A + int(A/4)
	JD := int(365.25*float64(year+4716)) + int(30.6001*float64(month+1)) + int(day) + B
	// Add fractional day based on time.
	dayFraction := (float64(t.Hour()) + float64(t.Minute())/60 + float64(t.Second())/3600 + float64(t.Nanosecond())/1e9/3600) / 24
	return float64(JD) + dayFraction - 1524.5
}
