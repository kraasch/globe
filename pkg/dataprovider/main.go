// Package dataprov is an abstraction from different ways to calculate the same data about moon, sun, etc.
package dataprov

import (
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/moonposition"

	// local packages.
	"github.com/kraasch/geo/pkg/astro"
)

// some variables.
var (
	// Go's simpleTimeLayout strings.
	simpleTimeLayout = "2006-01-02 15:04:05"
	// priciseTimeLayout = "2006-01-02 15:04:05.999999999" // NOTE: for reference.
)

///////////////////
// DATA PROVIDER //
///////////////////

type DataProvider struct {
	time time.Time
}

func (dp *DataProvider) SetTime(timeStr string) error {
	t, err := time.ParseInLocation(simpleTimeLayout, timeStr, time.UTC)
	if err != nil {
		return err
	}
	dp.time = t
	return nil
}

///////////////////////////
// GENERAL DATA PROVIDER //
///////////////////////////

type GeneralDataProvider struct {
	DataProvider
}

// #######################
// No. 1 -- geo/astro
// #######################

type GeoGeneralDataProvider struct {
	GeneralDataProvider
}

func NewGeoGeneralDataProvider() GeoGeneralDataProvider {
	return GeoGeneralDataProvider{}
}

func (p *GeoGeneralDataProvider) JulianDate() float64 {
	return astro.JulianDate(p.time)
}

// #######################
// No. 2 -- soniakeys/meeus/v3
// #######################

type KeysGeneralDataProvider struct {
	GeneralDataProvider
}

func NewKeysGeneralDataProvider() KeysGeneralDataProvider {
	return KeysGeneralDataProvider{}
}

func (p *KeysGeneralDataProvider) JulianDate() float64 {
	return julian.TimeToJD(p.time)
	// jd := julian.CalendarGregorianToJD(now.Year(), int(now.Month()), now.Day()) // NOTE: this also exists.
}

////////////////////////
// MOON DATA PROVIDER //
////////////////////////

type MoonDataProvider struct {
	DataProvider
}

// #######################
// No. 1 -- geo/astro // TODO: impelemnt.
// #######################

type GeoMoonDataProvider struct { // TODO: impelemnt.
	MoonDataProvider
}

func NewGeoMoonDataProvider() GeoMoonDataProvider { // TODO: impelemnt.
	return GeoMoonDataProvider{}
}

func (p *GeoMoonDataProvider) GeocentricCoords() (float64, float64) { // TODO: impelemnt.
	return 0.0, 0.0 // astro.MoonPosition(p.time)
}

// #######################
// No. 2 -- soniakeys/meeus/v3
// #######################

type KeysMoonDataProvider struct {
	MoonDataProvider
}

func NewKeysMoonDataProvider() KeysMoonDataProvider {
	return KeysMoonDataProvider{}
}

func (p *KeysMoonDataProvider) GeocentricCoords() (float64, float64) {
	jd := julian.TimeToJD(p.time)
	// NOTE: third return value of moonposition.Position() is the distance between earth and moon in km.
	lon, lat, dist := moonposition.Position(jd)
	fmt.Printf("MOON: lat: %f, lon: %f, dist: %f\n", lat, lon, dist)
	return float64(lat.Deg()), float64(lon.Deg() - 180.0)
}

// #######################
// No. 3 -- hablullah/go-sampa
// #######################

type SampaMoonDataProvider struct {
	MoonDataProvider
}

func NewSampaMoonDataProvider() SampaMoonDataProvider {
	return SampaMoonDataProvider{}
}

func (p *SampaMoonDataProvider) GeocentricCoords() (float64, float64) {
	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81}
	moonpos, _ := sampa.GetMoonPosition(p.time, jakarta, nil)
	return moonpos.GeocentricLatitude - 360.0, moonpos.GeocentricLongitude - 180.0
}

// // MAKE THIS INTO ANOTHER PROVIDER. // TODO: implement.
// 	mp "github.com/janczer/goMoonPhase"
// func MoonLon(date time.Time) float64 {
// 	phase := mp.New(date)
// 	return phase.Longitude() - 180.0
// }

///////////////////////
// SUN DATA PROVIDER //
///////////////////////

// TODO: implement sun data providers.
