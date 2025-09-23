// Package geomap is is used to show an ASCII world map on the CLI.
package geomap

import (
	"errors"
	"fmt"
	"strings"
)

const (
	markerSpot     = '▣'
	markerMoon     = '●'
	markerSun      = '☼'
	defaultSubLine = "                        " // 24 spaces.
)

var NL = fmt.Sprintln()

var MAP = "    _,--._  _._.--.--.._" + NL +
	"=.--'=_',-,:`;_      .,'" + NL +
	",-.  _.)  (``-;_   .'   " + NL +
	"   '-:_    `) ) .''=.   " + NL +
	"     ) )    ()'    ='   " + NL +
	"     |/            (_) =" + NL +
	"     -                  "

type World struct {
	name string // TODO: delete later.
}

type Subline struct {
	Line string
}

func NewWorld() World {
	return World{"default"}
}

func NewSubLine() Subline {
	return Subline{defaultSubLine}
}

// AddMoon adds a moon symbol.
func (s *Subline) AddMoon(lon float64) error {
	return s.Add(lon, markerMoon)
}

// AddSun adds a sun symbol.
func (s *Subline) AddSun(lon float64) error {
	return s.Add(lon, markerSun)
}

// Add adds a character into the marker line under the world map (which is of length 24).
func (s *Subline) Add(lon float64, marker rune) error {
	col, err0 := lon2col(lon)
	if err0 != nil {
		return err0
	}
	markedLine, reserr := markMap(col, 0, s.Line, marker)
	s.Line = markedLine
	return reserr
}

// lon2col turns a longitude between -180 and +180 into a value from 0 to 23.
func lon2col(lon float64) (int, error) {
	// "1-987654321 123456789+12"
	// "    _,--._  _._.--.--.._"
	// "=.--'=_',-,:`;_      .,'"
	// ",-.  _.)  (``-;_   .'   "
	// "▣  '-:_    `) ) .''=.   "
	// "     ) )    ()'    ='   "
	// "     |/            (_) ="
	// "     -                  "
	//  A          A           A
	//  |          |           |
	//  |          0           | ==> 360/24 = 15        // 24 columns, 360 degrees on the globe.
	//  |                      |                        // thus: 0-15 degrees are the first column, and so on.
	//  |                      | func lon2col( lon ):   // write a function to translate latitude to columns.
	//  .__ -180(E)  +180(W) __.   lon = lon+180        // start with 0 at the left.
	//                             return int(lon / 15) // divide without rest.
	if lon < -180 || lon > +180 { // this allows lon to be equal to +180.
		return 0, errors.New("longitude value out of bounds (lon < -180 || lon > +180)")
	}
	lon = lon + 180

	// // NOTE: optional, correct to the left.
	// // NOTE: implement later, maybe.
	// correction := -7.5 // correct 7.5 degrees to the left, in order to center around each column.
	// lon = lon + correction

	// divide into 24 zones (of 15 degrees each).
	res := int(lon / 15)

	// keep value in bounds in case lon was exactly +180.
	if res == 24 {
		res = 23
	}

	// return result.
	return res, nil
}

// lat2row turns a latitude between -90 and +90 into a value from 0 to 6.
func lat2row(lat float64) (int, error) {
	// "    _,--._  _._.--.--.._" // 0 for lat. +90 to +50. (= 40 degrees in total).  + ==> North.
	// "=.--'=_',-,:`;_      .,'" // 1 for lat. +50 to +35. (= 15 degrees in total).  - ==> South.
	// ",-.  _.)  (``-;_   .'   " // 2 for lat. +35 to +20. (= 15 degrees in total).
	// "   '-:_    `) ) .''=.   " // 3 for lat. +20 to -20. (= 40 degrees in total).
	// "     ) )    ()'    ='   " // 4 for lat. -20 to -35. (= 15 degrees in total).
	// "     |/            (_) =" // 5 for lat. -35 to -50. (= 15 degrees in total).
	// "     -                  " // 6 for lat. -50 to -90. (= 40 degrees in total).  all lat. degrees add up to 180.
	if lat < -90 || lat > +90 {
		return 0, errors.New("latitude value out of bounds (lat < -90 or lat > +90)")
	}
	row := 0
	switch {
	case lat <= +90 && lat > +50:
		row = 0
	case lat <= +50 && lat > +35:
		row = 1
	case lat <= +35 && lat > +20:
		row = 2
	case lat <= +20 && lat > -20:
		row = 3
	case lat <= -20 && lat > -35:
		row = 4
	case lat <= -35 && lat > -50:
		row = 5
	case lat <= -50 && lat >= -90: // NOTE: allows values to be from +90 to -90 including both borders.
		row = 6
	}
	return row, nil
}

// markMap replaces the character at (x, y) in the multi-line string MAP with a mark.
func markMap(x, y int, str string, marker rune) (string, error) {
	lines := strings.Split(str, "\n")
	if y < 0 || y >= len(lines) {
		return "", fmt.Errorf("y coordinate out of range")
	}
	line := lines[y]
	if x < 0 || x >= len(line) {
		return "", fmt.Errorf("x coordinate out of range")
	}
	// Replace character at position x
	runes := []rune(line)
	runes[x] = marker
	lines[y] = string(runes)

	return strings.Join(lines, "\n"), nil
}

// makeBox creates a box around some string and adds markers around the box.
func makeBox(lat, lon float64, str string) (string, error) { // TODO: implement.
	col, err0 := lon2col(lon)
	row, err1 := lat2row(lat)
	if err0 != nil {
		return "", err0
	}
	if err1 != nil {
		return "", err1
	}
	lines := strings.Split(str, "\n")
	lineLen := 24
	// Create top and bottom border
	topBorder := ""
	topBorder += "├" + strings.Repeat("─", col)
	topBorder += "▼"
	topBorder += strings.Repeat("─", lineLen-col-1) + "┤"
	bottomBorder := strings.Replace(topBorder, "▼", "▲", 1)
	var boxedLines []string
	for i, line := range lines {
		// Surround the line with │ and spaces
		boxedLine := ""
		if i == row {
			boxedLine = fmt.Sprintf("▶%s◀", line)
		} else {
			boxedLine = fmt.Sprintf("│%s│", line)
		}
		boxedLines = append(boxedLines, boxedLine)
	}
	// Combine all parts
	result := topBorder + "\n"
	result += strings.Join(boxedLines, "\n") + "\n"
	result += bottomBorder
	return result, nil
}

// PrintCoordBox uses PrintCoord and then creates a box around it.
func (w *World) PrintCoordBox(lat, lon float64) (string, error) {
	inner, err0 := w.PrintCoord(lat, lon)
	if err0 != nil {
		return "", err0
	}
	boxed, err1 := makeBox(lat, lon, inner)
	if err1 != nil {
		return "", err1
	}
	return boxed, nil
}

// PrintCoord calculates x and y coordinates within 2D string world map from latitude and longitude values and marks the location within the 2D string.
func (w *World) PrintCoord(lat, lon float64) (string, error) {
	col, err0 := lon2col(lon)
	row, err1 := lat2row(lat)
	if err0 != nil {
		return MAP, err0
	}
	if err1 != nil {
		return MAP, err1
	}
	markedMap, err2 := markMap(col, row, MAP, markerSpot)
	return markedMap, err2
}

func (w *World) PrintInner() string {
	return "    _,--._  _._.--.--.._" + NL +
		"=.--'=_',-,:`;_      .,'" + NL +
		",-.  _.)  (``-;_   .'   " + NL +
		"   '-:_    `) ) .''=.   " + NL +
		"     ) )    ()'    ='   " + NL +
		"     |/            (_) =" + NL +
		"     -                  "
}

func (w *World) PrintBlank() string {
	return "┌────────────────────────┐" + NL +
		"│1-987654321 123456789+12│" + NL +
		"├────────────────────────┤" + NL +
		"│    _,--._  _._.--.--.._│" + NL +
		"│=.--'=_',-,:`;_      .,'│" + NL +
		"│,-.  _.)  (``-;_   .'   │" + NL +
		"│   '-:_    `) ) .''=.   │" + NL +
		"│     ) )    ()'    ='   │" + NL +
		"│     |/            (_) =│" + NL +
		"│     -                  │" + NL +
		"├────────────────────────┤" + NL +
		"│                        │" + NL +
		"└────────────────────────┘"
}

func (w *World) PrintDemo() string {
	return "┌────────────────────────┐" + NL +
		"│1-987654321 123456789+12│" + NL +
		"├───────────▼────────────┤" + NL +
		"│    _,--._  _._.--.--.._│" + NL +
		"▶=.--'=_',-,▣`;_      .,'◀" + NL +
		"│,-.  _.)  (``-;_   .'   │" + NL +
		"│   '-:_    `) ) .''=.   │" + NL +
		"│     ) )    ()'    ='   │" + NL +
		"│     |/            (_) =│" + NL +
		"│     -                  │" + NL +
		"├───────────▲────────────┤" + NL +
		"│   ☼            ●       │" + NL +
		"└────────────────────────┘"
}
