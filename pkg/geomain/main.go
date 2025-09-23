// Package geomain is used to nicely print information calculated or retreiveed from geocalc and geoauto.
package geomain

import (
	geomap "github.com/kraasch/geo/pkg/geomap"
)

type GeoData struct {
	world geomap.World
}

func Toast() string {
	return "Toast!"
}

func New() GeoData {
	return GeoData{geomap.NewMarkedWorld(0.0, 0.0, 0.0, 0.0)}
}

func (gd *GeoData) UpdateData() {
	// TODO: implement.
	gd.world = geomap.NewMarkedWorld(0.0, 0.0, 0.0, 0.0)
}

func (gd *GeoData) PrintData() string {
	data, err := gd.world.Print()
	if err != nil {
		return err.Error()
	}
	return data
}
