// Package dataprov is an abstraction from different ways to calculate the same data about moon, sun, etc.
package dataprov

import (
	"time"

	"github.com/hablullah/go-sampa"
	osman "github.com/nathan-osman/go-sunrise"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/moonposition"

	// local packages.
	"github.com/kraasch/globe/pkg/astro"
)

// some variables.
var (
	// Go's simpleTimeLayout strings.
	simpleTimeLayout = "2006-01-02 15:04:05"
	// priciseTimeLayout = "2006-01-02 15:04:05.999999999" // NOTE: for reference.
)

/////////////////////////////////////////////////////////
//////////////////////////////////////// DATA PROVIDER //
/////////////////////////////////////////////////////////

type DataProvider struct {
	time time.Time
}

type DataProviderInterface interface {
	SetTime(timeStr string) error
}

func (dp *DataProvider) SetTime(timeStr string) error {
	t, err := time.ParseInLocation(simpleTimeLayout, timeStr, time.UTC)
	if err != nil {
		return err
	}
	dp.time = t
	return nil
}

/////////////////////////////////////////////////////////////////
//////////////////////////////////////// GENERAL DATA PROVIDER //
/////////////////////////////////////////////////////////////////

type GeneralDataProviderInterface interface {
	DataProviderInterface
	JulianDate() float64
}

// #######################
// No. 1 -- globe/astro
// #######################

type GlobeGeneralDataProvider struct {
	DataProvider
}

func (p *GlobeGeneralDataProvider) JulianDate() float64 {
	return astro.JulianDate(p.time)
}

// #######################
// No. 2 -- soniakeys/meeus/v3
// #######################

type KeysGeneralDataProvider struct {
	DataProvider
}

func (p *KeysGeneralDataProvider) JulianDate() float64 {
	return julian.TimeToJD(p.time)
	// jd := julian.CalendarGregorianToJD(now.Year(), int(now.Month()), now.Day()) // NOTE: this also exists.
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////// MOON DATA PROVIDER //
//////////////////////////////////////////////////////////////

type MoonDataProviderInterface interface {
	DataProviderInterface
	MoonsGeocentricCoords() (float64, float64)
}

// #######################
// No. 1 -- soniakeys/meeus/v3
// #######################

type KeysMoonDataProvider struct {
	DataProvider
}

func (p KeysMoonDataProvider) MoonsGeocentricCoords() (float64, float64) {
	jd := julian.TimeToJD(p.time)
	// NOTE: third return value of moonposition.Position() is the distance between earth and moon in km.
	lonAngle, latAngle, _ := moonposition.Position(jd)
	lat := float64(latAngle.Deg())
	lon := float64(lonAngle.Deg() - 180.0)
	return lat, lon
}

// #######################
// No. 2 -- hablullah/go-sampa
// #######################

type SampaMoonDataProvider struct {
	DataProvider
}

func (p SampaMoonDataProvider) MoonsGeocentricCoords() (float64, float64) {
	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81}
	moonpos, _ := sampa.GetMoonPosition(p.time, jakarta, nil)
	return moonpos.GeocentricLatitude - 360.0, moonpos.GeocentricLongitude - 180.0
}

// // MAKE THIS INTO ANOTHER PROVIDER. // TODO: implement.
// // #######################
// // No. 3 -- xxx
// // #######################
// 	mp "github.com/janczer/goMoonPhase"
// func MoonLon(date time.Time) float64 {
// 	phase := mp.New(date)
// 	return phase.Longitude() - 180.0
// }

/////////////////////////////////////////////////////////////
//////////////////////////////// SUNSET + SUNRISE PROVIDER //
/////////////////////////////////////////////////////////////

type SunsetSunriseDataProviderInterface interface {
	DataProviderInterface
	Sunset(float64, float64) time.Time
	Sunrise(float64, float64) time.Time
}

// #######################
// No. 1 -- globe/astro
// #######################

type OsmanSunsetSunriseDataProvider struct {
	DataProvider
}

func (p OsmanSunsetSunriseDataProvider) Sunrise(lat, lon float64) time.Time {
	date := p.time
	sr, _ := osman.SunriseSunset(lat, lon, date.Year(), date.Month(), date.Day())
	return sr
}

func (p OsmanSunsetSunriseDataProvider) Sunset(lat, lon float64) time.Time {
	date := p.time
	_, ss := osman.SunriseSunset(lat, lon, date.Year(), date.Month(), date.Day())
	return ss
}

/////////////////////////////////////////////////////////////
//////////////////////////////////////// SUN DATA PROVIDER //
/////////////////////////////////////////////////////////////

type SunDataProviderInterface interface {
	DataProviderInterface
	SunsGeocentricCoords() (float64, float64)
}

// #######################
// No. 1 -- globe/astro
// #######################

type GlobeSunDataProvider struct {
	DataProvider
}

func (p GlobeSunDataProvider) SunsGeocentricCoords() (float64, float64) {
	return astro.SunsGeocentricCoords(p.time)
}

// // MAKE THIS INTO ANOTHER PROVIDER. // TODO: implement.
// // #######################
// // No. 2 -- soniakeys/meeus, ie. v3/solar package.
// // #######################

// // MAKE THIS INTO ANOTHER PROVIDER. // TODO: implement.
// // #######################
// // No. 3 -- xxx
// // #######################
// 	solar "github.com/observerly/sidera/pkg/solar"
// func SunLat(date time.Time) float64 { // TODO: fix.
// 	equatorialCoord := solar.GetEquatorialCoordinate(date)
// 	return equatorialCoord.Declination
// }

// // MAKE THIS INTO ANOTHER PROVIDER. // TODO: implement.
// // #######################
// // No. 4 -- xxx
// // #######################
// func SunLon(date time.Time) float64 { // TODO: fix.
// 	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81} // NOTE: just somewhere on this planet, to kick off a calculation with much overhead.
// 	sunpos, _ := sampa.GetSunPosition(date, jakarta, nil)
// 	sl := sunpos.GeocentricLongitude
// 	return sl - 180.0
// }
