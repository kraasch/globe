// Package geoauto can get the location or time zone of a user machine automatically. This package is also not easily testable.
package geoauto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	ctor "github.com/kraasch/goconf/pkg/configurator"
)

const (
	configName  = "renamer.toml"
	configPath  = ".config/geo/"
	defaultData = ""
)

var config = ctor.Configurator{
	ConfigFileName: configName,
	PathToConfig:   configPath,
	DefaultConfig:  defaultData,
}

// GeoLocation struct to parse JSON response from IP geolocation API.
type GeoLocation struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

// bufferWebLocalization looks if lon+lat have been stored recently in ~/.local/geo/data.txt
// if yes, it uses those values, otherwise it tries to retreive new values from the web, if that worked it writes them into the file.
func bufferWebLocalization() (float64, float64, error) {
	// rawData := config.AutoReadConfig()
	// TODO: implement.
	return 0.1, 0.1, nil
}

// complicatedWebLocalization gets user's location based on IP.
func complicatedWebLocalization() (float64, float64, error) {
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
	loc, err0 := SimpleSystemLocalization()
	lon, lat, err1 := complicatedWebLocalization()
	result := ""
	if err0 != nil || err1 != nil {
		result = "location not found"
	} else {
		result = fmt.Sprintf("zone: %s, lon: %.2f, lat: %.2f", loc.String(), lon, lat)
	}
	return result
}
