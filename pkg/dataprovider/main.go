// Package dataprov is an abstraction from different ways to calculate the same data about moon, sun, etc.
package dataprov

import (
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/soniakeys/meeus/v3/julian"

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
// No. 2 -- soniakeys/meees/v3
// #######################

type KeysGeneralDataProvider struct {
	GeneralDataProvider
}

func NewKeysGeneralDataProvider() KeysGeneralDataProvider {
	return KeysGeneralDataProvider{}
}

func (p *KeysGeneralDataProvider) JulianDate() float64 {
	return julian.TimeToJD(p.time)
}

////////////////////////
// MOON DATA PROVIDER //
////////////////////////

type MoonDataProvider struct {
	DataProvider
}

// #######################
// No. 1 -- geo/astro
// #######################

type GeoMoonDataProvider struct {
	MoonDataProvider
}

func NewGeoMoonDataProvider() GeoMoonDataProvider {
	return GeoMoonDataProvider{}
}

func (p *GeoMoonDataProvider) GeocentricCoords() (float64, float64) {
	return astro.MoonPosition(p.time)
}

// #######################
// No. 2 -- hablullah/go-sampa
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
