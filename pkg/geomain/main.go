// Package geomain is used to nicely print information calculated or retreiveed from geocalc and geoauto.
package geomain

import (
	"fmt"
	"time"

	geocalc "github.com/kraasch/geo/pkg/geocalc"
	geomap "github.com/kraasch/geo/pkg/geomap"
)

type GeoData struct {
	world geomap.World
	time  time.Time
}

func Toast() string {
	return "Toast!"
}

func New() GeoData {
	return GeoData{}
}

func (gd *GeoData) UpdateData() {
	// TODO: implement.
	gd.world = geomap.NewMarkedWorld(0.0, 0.0, 0.0, 0.0)
	gd.time = time.Now()
}

func (gd *GeoData) PrintData() string {
	NL := fmt.Sprintln()
	data, err := gd.world.Print()
	moon := geocalc.MoonPhase(gd.time)
	sun := geocalc.SunRiseAndSet(0.0, 0.0, gd.time) // TODO: insert lon and lat from gd.world.
	if err != nil {
		return err.Error()
	}
	return data + NL + moon + NL + sun
}
