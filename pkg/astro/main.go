// Package astro is rolling geo's own calculations.
package astro

import (
	"math"
	"time"
)

// ###################################################################
// ## GENERAL DATA
// ###################################################################

func JulianDate(t time.Time) float64 {
	// JulianDate should only take a UTC time. // TODO: check for all dates to be UTC.
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

func roughDeltaTAround21stCentury(year int) float64 { // TODO: make more precise.
	t := float64(year-2000) / 100.0
	deltaT := 102.0 + 102.0*t + 25.3*t*t
	deltaT += 0.37 * float64(year-2100)
	return deltaT
}

func JulianEphemerisDay(jd float64, t time.Time) float64 { // TODO: test.
	dt := roughDeltaTAround21stCentury(t.Year())
	return jd + dt/86400.0
}

// ###################################################################
// ## MOON POSITION
// ###################################################################

/*
// TODO: implement.
func MoonPosition(t time.Time) (longitude, latitude float64) { // NOTE: this is garbage.
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
*/

// ###################################################################
// ## SUNSET AND SUNRISE
// ###################################################################

// func Sunset(t time.Time, lat float64, lon float64) time.Time {
// 	// Convert the date to Julian Day
// 	year, month, day := t.Date()
// 	// Calculate the approximate Julian Day
// 	JD := julianDay(year, int(month), day) // TODO: remove this function
// 	// Calculate days since J2000.0
// 	n := JD - 2451545.0 + 0.0008
// 	// Mean solar noon
// 	Jstar := n - lon/360.0
// 	// Solar mean anomaly
// 	tempM := 357.5291 + 0.98560028*Jstar
// 	M := math.Mod(tempM, 360)
// 	// Convert to radians
// 	M_rad := M * math.Pi / 180
// 	// Equation of center
// 	C := (1.9148*math.Sin(M_rad) + 0.0200*math.Sin(2*M_rad) + 0.0003*math.Sin(3*M_rad))
// 	// Ecliptic longitude
// 	tempLambda := M + C + 180 + 102.9372
// 	λ := math.Mod(tempLambda, 360)
// 	// Convert to radians
// 	λ_rad := λ * math.Pi / 180
// 	// Sun declination
// 	sinδ := math.Sin(λ_rad) * math.Sin(23.45*math.Pi/180)
// 	// Calculate the hour angle
// 	lat_rad := lat * math.Pi / 180
// 	cosH := (math.Cos(math.Pi/180*90.833) - math.Sin(lat_rad)*sinδ) / (math.Cos(lat_rad) * math.Cos(asin(sinδ)))
// 	if cosH > 1 {
// 		// Sun always below horizon
// 		return time.Time{} // no sunset
// 	}
// 	if cosH < -1 {
// 		// Sun always above horizon
// 		return time.Time{} // no sunset
// 	}
// 	H := math.Acos(cosH) // in radians
// 	// Convert hour angle to hours
// 	H_deg := H * 180 / math.Pi
// 	// Calculate sunset time in UTC
// 	// Solar noon in UTC
// 	Jt := Jstar + H_deg/360
// 	// Convert Julian day back to time
// 	sunsetUTC := julianToTime(Jt)
// 	return sunsetUTC
// }
//
// func julianDay(year, month, day int) float64 { // TODO: remove.
// 	if month <= 2 {
// 		year -= 1
// 		month += 12
// 	}
// 	A := math.Floor(float64(year) / 100)
// 	B := 2 - A + math.Floor(A/4)
// 	JD := math.Floor(365.25*float64(year+4716)) + math.Floor(30.6001*float64(month+1)) + float64(day) + B - 1524.5
// 	return JD
// }
//
// func julianToTime(J float64) time.Time {
// 	// Convert Julian Day to date
// 	J += 0.5
// 	Z := math.Floor(J)
// 	F := J - Z
// 	var A float64
// 	if Z >= 2299161 {
// 		alpha := math.Floor((Z - 1867216.25) / 36524.25)
// 		A = Z + 1 + alpha - math.Floor(alpha/4)
// 	} else {
// 		A = Z
// 	}
// 	B := A + 1524
// 	C := math.Floor((B - 122.1) / 365.25)
// 	D := math.Floor(365.25 * C)
// 	E := math.Floor((B - D) / 30.6001)
//
// 	day := B - D - math.Floor(30.6001*E) + F
// 	month := E - 1
// 	if month > 12 {
// 		month -= 12
// 	}
// 	year := C - 4716
// 	hourFraction := day - math.Floor(day)
// 	hour := int(hourFraction * 24)
// 	minute := int((hourFraction*24 - float64(hour)) * 60)
// 	second := int((((hourFraction*24 - float64(hour)) * 60) - float64(minute)) * 60)
//
// 	return time.Date(int(year), time.Month(int(month)), int(math.Floor(day)), hour, minute, second, 0, time.UTC)
// }
//
// func asin(x float64) float64 {
// 	return math.Asin(x)
// }
//
// func Sunrise(t time.Time, lat float64, lon float64) time.Time {
//   return t // TODO: implement.
// }

// ###################################################################
// ## SUN POSITION
// ###################################################################

// Constants
const (
	deg2rad = math.Pi / 180
	rad2deg = 180 / math.Pi
)

// Helper functions
func sinDeg(deg float64) float64 {
	return math.Sin(deg2rad * deg)
}

func cosDeg(deg float64) float64 {
	return math.Cos(deg2rad * deg)
}

// Calculate number of days since J2000.0
func daysSinceJ2000(jd float64) float64 {
	return jd - 2451545.0
}

// Compute the Sun's mean longitude (degrees)
func meanLongitude(days float64) float64 {
	L0 := 280.46646 + 0.9856474*days
	return math.Mod(L0, 360)
}

// Compute the Sun's mean anomaly
func meanAnomaly(days float64) float64 {
	M := 357.52911 + 0.98560028*days
	return math.Mod(M, 360)
}

// Compute the Sun's ecliptic longitude (degrees)
func eclipticLongitude(L, M float64) float64 {
	C := (1.9148 * sinDeg(M)) + (0.0200 * sinDeg(2*M)) + (0.0003 * sinDeg(3*M))
	λ := L + C
	return math.Mod(λ, 360)
}

// Compute the Sun's declination (δ)
func sunDeclination(λ float64) float64 {
	ε := 23.4397 // obliquity of the ecliptic
	sinδ := sinDeg(ε) * sinDeg(λ)
	δ := math.Asin(sinδ)
	return δ * rad2deg
}

// Compute the right ascension (α) (degrees)
func rightAscension(λ float64) float64 {
	ε := 23.4397
	sinλ := sinDeg(λ)
	cosλ := cosDeg(λ)
	tanα := cosDeg(ε) * sinλ / cosλ
	α := math.Atan(tanα) * rad2deg
	if cosλ < 0 {
		α += 180
	}
	return math.Mod(α+360, 360)
}

// Calculate sidereal time at Greenwich (degrees)
func siderealTime(jd float64, longitude float64) float64 {
	T := (jd - 2451545.0) / 36525
	S := 280.46061837 + 360.98564736629*(jd-2451545) + 0.000387933*T*T - T*T*T/38710000
	S = math.Mod(S, 360)
	if S < 0 {
		S += 360
	}
	// Local sidereal time
	lst := S + longitude
	return math.Mod(lst, 360)
}

func SunsGeocentricCoords(t time.Time) (float64, float64) {
	jd := JulianDate(t)
	days := daysSinceJ2000(jd)
	// Compute Sun's position
	L := meanLongitude(days)
	M := meanAnomaly(days)
	λ := eclipticLongitude(L, M)
	// Sun's declination (subsolar latitude)
	δ := sunDeclination(λ)
	// Compute the Sun's right ascension
	α := rightAscension(λ)
	// Compute sidereal time at Greenwich
	θ := siderealTime(jd, 0)
	// Calculate hour angle (degrees)
	H := θ - α
	H = math.Mod(H+360, 360)
	if H > 180 {
		H -= 360
	}
	// Subsolar latitude is declination
	subsolarLat := δ
	// Subsolar longitude (degrees)
	// Local solar time (hours)
	lstHours := H / 15.0
	subsolarLon := lstHours * 15.0
	// Adjust to [-180, 180]
	if subsolarLon > 180 {
		subsolarLon -= 360
	}
	// fmt.Printf("Sun's Apparent Longitude: %.2f°\n", λ)
	return subsolarLat, subsolarLon
}
