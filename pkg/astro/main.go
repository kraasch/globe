// Package astro is rolling geo's own calculations.
package astro

import (
	"math"
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

func MoonPosition(t time.Time) (longitude, latitude float64) {
	jd := JulianDate(t)
	D := jd - 2451545.0                         // Number of days since J2000.0
	L0 := math.Mod(218.316+13.176396*D, 360)    // Mean longitude of the Moon
	Dm := math.Mod(297.850192+12.190749*D, 360) // Mean elongation of the Moon
	// Ms := math.Mod(357.529109+0.98560028*D, 360) // Sun's mean anomaly // TODO: remove later.
	Mm := math.Mod(134.963+13.064993*D, 360) // Moon's mean anomaly
	// Convert to radians.
	// L0r := L0 * math.Pi / 180 // TODO: remove later.
	Dmr := Dm * math.Pi / 180
	// Msr := Ms * math.Pi / 180 // TODO: remove later.
	Mmr := Mm * math.Pi / 180
	// Simplified ecliptic longitude calculation
	longitude = L0 +
		6.289*math.Sin(Mmr) + // Moon's anomaly
		1.274*math.Sin(2*Dmr-Mmr) +
		0.658*math.Sin(2*Dmr) +
		0.214*math.Sin(2*Mmr) +
		0.110*math.Sin(Dmr)
	// Normalize to [0, 360)
	longitude = math.Mod(longitude, 360)
	// For latitude, a simplified approach:
	latitude = 5.128 * math.Sin(Mmr+0.5*Dmr) // approximate lunar latitude
	return longitude, latitude
}
