// Package dataprov is an abstraction from different ways to calculate the same data about moon, sun, etc.
package dataprov

import (
	"time"

	"github.com/hablullah/go-sampa"
)

// some variables.
var (
	// Go's simpleTimeLayout strings.
	simpleTimeLayout = "2006-01-02 15:04:05"
	// priciseTimeLayout = "2006-01-02 15:04:05.999999999" // NOTE: for reference.
)

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

func (sp *SampaMoonDataProvider) GeocentricCoords() (float64, float64) {
	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81}
	moonpos, _ := sampa.GetMoonPosition(sp.time, jakarta, nil)
	return moonpos.GeocentricLatitude - 360.0, moonpos.GeocentricLongitude - 180.0
}

// // MAKE THIS INTO ANOTHER PROVIDER.
// 	mp "github.com/janczer/goMoonPhase"
// func MoonLon(date time.Time) float64 { // TODO: fix.
// 	phase := mp.New(date)
// 	return phase.Longitude() - 180.0
// }
