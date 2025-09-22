// Package geoauto can get the location or time zone of a user machine automatically. This package is also not easily testable.
package geoauto

import (
	"fmt"
	"time"
)

// GetCurrentLocation returns values like "America/New_York" or "UTC", by asking the system.
func GetCurrentLocation() (*time.Location, error) {
	loc, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func Toast() string {
	loc, err := GetCurrentLocation()
	result := ""
	if err != nil {
		result = "location not found"
	} else {
		result = fmt.Sprintf("%#v\n", loc.String())
	}
	return result
}
