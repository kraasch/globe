// Package geoview is used to nicely print information calculated or retreiveed from geocalc and geoauto.
package geoview

import (
	"fmt"
	"strings"
	"time"

	geocalc "github.com/kraasch/geo/pkg/geocalc"
	geomap "github.com/kraasch/geo/pkg/geomap"
	geoweb "github.com/kraasch/geo/pkg/geoweb"
)

const (
	B0 = "\x1b[1;38;2;0;0;0m"       // ANSI foreground color (= black).
	A0 = "\x1b[1;38;2;10;10;10m"    // ANSI foreground color (= gray 0).
	A1 = "\x1b[1;38;2;100;100;100m" // ANSI foreground color (= gray 1).
	A2 = "\x1b[1;38;2;150;150;150m" // ANSI foreground color (= gray 2).
	A3 = "\x1b[1;38;2;200;200;200m" // ANSI foreground color (= gray 3).
	W0 = "\x1b[1;38;2;255;255;255m" // ANSI foreground color (= white).
	R1 = "\x1b[1;38;2;255;0;0m"     // ANSI foreground color (= red).
	G1 = "\x1b[1;38;2;0;255;0m"     // ANSI foreground color (= green).
	B1 = "\x1b[1;38;2;0;0;255m"     // ANSI foreground color (= blue).
	O1 = "\x1b[1;38;2;255;150;0m"   // ANSI foreground color (= orange).
	Y1 = "\x1b[1;38;2;255;255;0m"   // ANSI foreground color (= yellow).
	N0 = "\x1b[0m"                  // ANSI clear formatting.
	P9 = "\x1b[48;5;56m"            // ANSI background color (= purple). // TODO: remove later.
	// B2 = "\x1b[1;38;2;100;100;100m" // ANSI foreground color (= gray). // TODO: remove later.
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
	gd.world = geomap.NewWorld()
	gd.world.Lat = 0.0
	gd.world.Lon = 0.0
	gd.world.MoonLon = 0.0
	gd.world.SunLon = 0.0
	gd.time = time.Now()
}

func (gd *GeoData) PrintDataVertically() string {
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

func surround(str, find, prefix, suffix string) string {
	return strings.ReplaceAll(str, find, prefix+find+suffix)
}

func colorizeSymbols(in string) string {
	in = surround(in, "▣", O1, N0)
	in = surround(in, "□", O1, N0)
	in = surround(in, "🜃", B1, N0)
	in = surround(in, "☼", R1, N0)
	in = surround(in, "○", A3, N0)
	// in = surround(in, "▼", A0, N0)
	// in = surround(in, "◀", A0, N0)
	// in = surround(in, "▲", A0, N0)
	// in = surround(in, "▶", A0, N0)
	return in
}

func (gd *GeoData) PrintInfo() string {
	NL := fmt.Sprintln()
	where := geoweb.LatAndLonAndTz()
	moon := geocalc.MoonPhase(gd.time)
	utc := geocalc.LocalAndUtcTime()
	sun := geocalc.SunRiseAndSet(0.0, 0.0, gd.time) // TODO: insert lon and lat from gd.world.
	return where + utc + NL + sun + NL + moon
}

func (gd *GeoData) PrintDataHorizontally() string {
	world := gd.PrintWorld()
	info := gd.PrintInfo()
	str := concatenateHorizontally(world, info)
	// TODO: implement color flag and check for it here.
	str = colorizeSymbols(str)
	return str
}
