// Package astro is rolling geo's own calculations.
package astro

import (
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

func JulianDate(t time.Time) float64 {
	// TODO: write my own function here.
	return julian.TimeToJD(t)
}
