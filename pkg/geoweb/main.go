// Package geoweb can get the location or time zone of a user machine from the web.
package geoweb

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	ctor "github.com/kraasch/goconf/pkg/configurator"
)

const (
	configName  = "geo.toml"
	configPath  = ".config/geo/"
	defaultData = ""
)

func NewWebBuffer() WebBuffer {
	return WebBuffer{}
}

type WebBuffer struct {
	Lat         float64
	Lon         float64
	LastRequest time.Time // TODO: use this in order not to make too many web requests.
}

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

func (wb *WebBuffer) Update() {
	wb.Lat, wb.Lon, _ = complicatedWebLocalization()
}

func (wb *WebBuffer) GetCoords() (float64, float64) {
	wb.Update() // TODO: remove this from here.
	return wb.Lat, wb.Lon
}

// bufferWebLocalization looks if lon+lat have been stored recently in ~/.local/geo/data.txt
// if yes, it uses those values, otherwise it tries to retreive new values from the web, if that worked it writes them into the file.
func bufferWebLocalization() (float64, float64, error) {
	// rawData := config.AutoReadConfig()
	// TODO: implement.
	return 0.1, 0.1, nil
}

// complicatedWebLocalization gets user's location based on IP.
func complicatedWebLocalization() (lat float64, lon float64, err error) {
	// lat = -41.48
	// lon = 120.22
	// err = nil
	// return lat, lon, err // TODO: use the real thing later when buffering is implemented.
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
