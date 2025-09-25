// Package web can get the location or time zone of a user machine from the web.
package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	ctor "github.com/kraasch/goconf/pkg/configurator"
	tzf "github.com/ringsaturn/tzf"
)

const (
	configName  = "geo.toml"
	configPath  = ".config/geo/"
	defaultData = ""
)

var config = ctor.Configurator{
	ConfigFileName: configName,
	PathToConfig:   configPath,
	DefaultConfig:  defaultData,
}

var tzFinder tzf.F

func initTzf() {
	if tzFinder == nil {
		var err error
		tzFinder, err = tzf.NewDefaultFinder()
		if err != nil {
			panic(err)
		}
	}
}

// GeoLocation struct to parse JSON response from IP geolocation API.
type GeoLocation struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func ConvertLatAndLonToTimezone(lat, lon float64) string {
	initTzf()                                // init tzFinder variable.
	tz := tzFinder.GetTimezoneName(lon, lat) // NOTE: Takes longitude-latitude order.
	return tz
}

// bufferWebLocalization looks if lon+lat have been stored recently in ~/.local/geo/data.txt
// if yes, it uses those values, otherwise it tries to retreive new values from the web, if that worked it writes them into the file.
func bufferWebLocalization() (float64, float64, error) {
	// rawData := config.AutoReadConfig()
	// TODO: implement.
	return 0.1, 0.1, nil
}

func LatAndLon() string {
	lat, lon, _ := complicatedWebLocalization()
	return fmt.Sprintf(" ▣ lat+lon: %.2f, %.2f", lat, lon)
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

func LatAndLonAndTz() string {
	lat, lon, _ := complicatedWebLocalization()
	tz := ConvertLatAndLonToTimezone(lat, lon)
	now := time.Now()
	utcOffset := timeToUtcOffset(now)
	return fmt.Sprintf(" □ lat+lon: %.2f, %.2f\n □ zone:    %s\n □ offset:  %s\n", lat, lon, tz, utcOffset)
}

// complicatedWebLocalization gets user's location based on IP.
func complicatedWebLocalization() (float64, float64, error) {
	return 53.48, 10.22, nil // TODO: use the real thing later when buffering is implemented.
	// Use an IP geolocation API (e.g., ip-api.com)
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	var result struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, 0, err
	}
	return result.Lat, result.Lon, nil
}

// SimpleSystemLocalization returns values like "America/New_York" or "UTC", by asking the system.
func SimpleSystemLocalization() (*time.Location, error) {
	loc, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func Toast() string { // TODO: implement better tests and functions in this package.
	// loc, err0 := SimpleSystemLocalization()
	// lon, lat, err1 := complicatedWebLocalization()
	// result := ""
	// if err0 != nil || err1 != nil {
	// 	result = "location not found"
	// } else {
	// 	result = fmt.Sprintf("zone: %s, lon: %.2f, lat: %.2f", loc.String(), lon, lat)
	// }
	// return result
	return "Toast!"
}
