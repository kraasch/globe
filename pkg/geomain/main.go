// Package geomain is used to nicely print information calculated or retreiveed from geocalc and geoauto.
package geomain

import (
	"fmt"
	"strings"
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

// concatenateHorizontally concatenates two multi-line strings horizontally
func concatenateHorizontally(str1, str2 string) string {
	// Split the input strings into slices of strings (lines)
	lines1 := strings.Split(str1, "\n")
	lines2 := strings.Split(str2, "\n")
	// Determine the maximum number of lines in both input strings
	maxLines := max(len(lines2), len(lines1))
	// Prepare an output slice of strings
	var result []string
	// Iterate through each line, combining lines horizontally
	for i := range maxLines {
		// Ensure we don't exceed the length of either slice by checking bounds
		line1 := ""
		line2 := ""
		if i < len(lines1) {
			line1 = lines1[i]
		}
		if i < len(lines2) {
			line2 = lines2[i]
		}
		// Concatenate the lines horizontally
		result = append(result, line1+line2)
	}
	// Join the result lines into a single string, separated by newlines
	return strings.Join(result, "\n")
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
	utc := geocalc.LocalAndUtcTime()
	sun := geocalc.SunRiseAndSet(0.0, 0.0, gd.time) // TODO: insert lon and lat from gd.world.
	if err != nil {
		return err.Error()
	}
	return data + NL + moon + NL + utc + NL + sun
}

func (gd *GeoData) PrintWorld() string {
	data, err := gd.world.Print()
	if err != nil {
		return err.Error()
	}
	return data
}

func (gd *GeoData) PrintInfo() string {
	NL := fmt.Sprintln()
	moon := geocalc.MoonPhase(gd.time)
	utc := geocalc.LocalAndUtcTime()
	sun := geocalc.SunRiseAndSet(0.0, 0.0, gd.time) // TODO: insert lon and lat from gd.world.
	return utc + NL + sun + NL + moon
}

func (gd *GeoData) PrintHorizontally() string {
	world := gd.PrintWorld()
	info := gd.PrintInfo()
	str := concatenateHorizontally(world, info)
	return str
}
