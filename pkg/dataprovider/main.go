// Package dataprov is an abstraction from different ways to calculate the same data about moon, sun, etc.
package dataprov

import (
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/soniakeys/meeus/v3/julian"
)

// some variables.
var (
	// Go's simpleTimeLayout strings.
	simpleTimeLayout = "2006-01-02 15:04:05"
	// priciseTimeLayout = "2006-01-02 15:04:05.999999999" // NOTE: for reference.
)

///////////////////////////
// GENERAL DATA PROVIDER //
///////////////////////////

type GeneralDataProvider struct {
	time time.Time
}

type KeysGeneralDataProvider struct {
	GeneralDataProvider
}

func NewKeysGeneralDataProvider() KeysGeneralDataProvider {
	return KeysGeneralDataProvider{}
}

func (gdp *GeneralDataProvider) SetTime(timeStr string) error {
	t, err := time.Parse(simpleTimeLayout, timeStr)
	if err != nil {
		return err
	}
	gdp.time = t
	return nil
}

// #######################
// No. 1 -- soniakeys/meees/v3
// #######################

func (kp *KeysGeneralDataProvider) JulianDate() float64 {
	return julian.TimeToJD(kp.time)
}

////////////////////////
// MOON DATA PROVIDER //
////////////////////////

type MoonDataProvider struct {
	time time.Time
}

type SampaMoonDataProvider struct {
	MoonDataProvider
}

func NewSampaMoonDataProvider() SampaMoonDataProvider {
	return SampaMoonDataProvider{}
}

func (mdp *MoonDataProvider) SetTime(timeStr string) error {
	t, err := time.Parse(simpleTimeLayout, timeStr)
	if err != nil {
		return err
	}
	mdp.time = t
	return nil
}

// #######################
// No. 1 -- hablullah/go-sampa
// #######################

func (sp *SampaMoonDataProvider) GeocentricCoords() (float64, float64) {
	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81}
	moonpos, _ := sampa.GetMoonPosition(sp.time, jakarta, nil)
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
