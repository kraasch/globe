// Package geomap is is used to show an ASCII world map on the CLI.
package geomap

import (
	"errors"
	"fmt"
	"strings"
)

const (
	markerSpot       = '▣'
	markerMoon       = '●'
	markerSun        = '☼'
	defaultSubLine   = "                        " // 24 spaces.
	defaultSidebar   = " \n \n \n \n \n \n "      // 7 spaces.
	defaultEmptySide = "   \n   \n   \n   \n   \n   \n   \n   \n   \n   \n   \n   \n   "
	div              = "│"
	top              = "┌────────────────────────┐"
	num              = "│1-987654321 123456789+12│"
	bot              = "└────────────────────────┘"
	padding          = "  "
	sidetop          = "┌──"           // without padding.
	sidebot          = "└──"           // without padding.
	sidetopPadded    = "   \n   \n┌──" // with padding.
	sidebotPadded    = "└──\n   \n   " // with padding.
	sideline         = "│\n│\n│\n│\n│\n│\n│"
	cornerBL         = "└"
	cornerBR         = "┘"
	cornerTR         = "┐"
	cornerTL         = "┌"
	topMark          = "▼"
	botMark          = "▲"
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
	Padded     bool
	Inactive   bool
	ShowSide   bool
	ShowTop    bool
	ShowBot    bool
	ShowAsMini bool // TODO: remove this field.
	Lat        float64
	Lon        float64
	MoonLat    float64
	MoonLon    float64
	SunLat     float64
	SunLon     float64
}

type Sidebar struct {
	Bar string
}

type Subline struct {
	Line string
}

func NewWorld() World {
	return World{}
}

func NewSidebar() Sidebar {
	return Sidebar{defaultSidebar}
}

func NewSubLine() Subline {
	return Subline{defaultSubLine}
}

func (s *Sidebar) AddMoon(lat float64, skipMark bool) error {
	return s.Add(lat, markerMoon, skipMark)
}

func (s *Sidebar) AddSun(lat float64, skipMark bool) error {
	return s.Add(lat, markerSun, skipMark)
}

// AddMoon adds a moon symbol.
func (s *Subline) AddMoon(lon float64, skipMark bool) error {
	return s.Add(lon, markerMoon, skipMark)
}

// AddSun adds a sun symbol.
func (s *Subline) AddSun(lon float64, skipMark bool) error {
	return s.Add(lon, markerSun, skipMark)
}

// Add adds a character into the marker line under the world map (which is of length 24).
func (s *Subline) Add(lon float64, marker rune, skipMark bool) error {
	col, err0 := lon2col(lon)
	if err0 != nil {
		return err0
	}
	markedLine, reserr := markMap(col, 0, s.Line, marker, skipMark)
	s.Line = markedLine
	return reserr
}

func (s *Sidebar) Add(lat float64, marker rune, skipMark bool) error {
	row, err0 := lat2row(lat)
	if err0 != nil {
		return err0
	}
	markedBar, reserr := markMap(0, row, s.Bar, marker, skipMark)
	s.Bar = markedBar
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
func markMap(x, y int, str string, marker rune, markerSkip bool) (string, error) {
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
	if !markerSkip {
		runes[x] = marker
	}
	lines[y] = string(runes)

	return strings.Join(lines, "\n"), nil
}

// makeBox creates a box around some string and adds markers around the box.
func makeBox(lat, lon float64, str string, skipMark bool) (string, error) { // TODO: implement.
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
	bottomBorder := ""
	topBorder += cornerTL + strings.Repeat("─", col)
	bottomBorder += cornerBL + strings.Repeat("─", col)
	if skipMark {
		topBorder += "─"
		bottomBorder += "─"
	} else {
		topBorder += topMark
		bottomBorder += botMark
	}
	topBorder += strings.Repeat("─", lineLen-col-1) + cornerTR
	bottomBorder += strings.Repeat("─", lineLen-col-1) + cornerBR
	var boxedLines []string
	for i, line := range lines {
		// Surround the line with │ and spaces
		boxedLine := ""
		if i == row {
			if skipMark {
				boxedLine = fmt.Sprintf("│%s│", line)
			} else {
				boxedLine = fmt.Sprintf("▶%s◀", line)
			}
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

// Print returns a string of the ASCII world data with its current state (defined in the struct variables).
// TODO: do major refactor in this entire function.
// TODO: implement errors.
func (w *World) Print() (string, error) {
	if w.ShowAsMini { // TODO: remove this flag from the struct entirely.
		box, _ := w.PrintCoordBox()
		return box, nil
	}
	// create the main box.
	box, _ := w.PrintCoordBox()
	// create a padded top.
	ttt := ""
	if w.ShowTop {
		ttt = top + NL + num + NL
	} else {
		if w.Padded {
			ttt = defaultSubLine + padding + NL + defaultSubLine + padding + NL
		} else {
			ttt = ""
		}
	}
	// create a padded bot.
	bbb := ""
	if w.ShowBot {
		line := NewSubLine()
		_ = line.AddMoon(w.MoonLon, w.Inactive)
		_ = line.AddSun(w.SunLon, w.Inactive)
		sub := line.Line
		bbb = div + sub + div + NL + bot
	} else {
		if w.Padded {
			bbb = defaultSubLine + padding + NL + defaultSubLine + padding
		} else {
			bbb = ""
		}
	}
	// create a padded side.
	sss := ""
	if w.ShowSide {
		side := NewSidebar()
		_ = side.AddMoon(w.MoonLat, w.Inactive)
		_ = side.AddSun(w.SunLat, w.Inactive)
		bar := side.Bar
		bar = ConcatenateHorizontally(sideline, bar)
		bar = ConcatenateHorizontally(bar, defaultSidebar)
		if w.Padded {
			sss = sidetopPadded + NL + bar + NL + sidebotPadded
		} else {
			// NOTE: if there is no sidebar padding is still needed,
			// but if there is not bot or top, padding must not be inserted.
			if w.ShowBot && w.ShowTop {
				sss = sidetopPadded + NL + bar + NL + sidebotPadded // as if padded.
			} else if w.ShowBot && !w.ShowTop {
				sss = sidetop + NL + bar + NL + sidebotPadded
			} else if !w.ShowBot && w.ShowTop {
				sss = sidetopPadded + NL + bar + NL + sidebot
			} else { // not padded + no bars => no padding.
				sss = sidetop + NL + bar + NL + sidebot
			}
		}
	} else {
		if w.Padded {
			sss = defaultEmptySide
		} else {
			sss = ""
		}
	}
	// add all together.
	res := ttt + box + NL + bbb             // add top and bot.
	res = ConcatenateHorizontally(sss, res) // add sidebar.
	return res, nil
}

// PrintCoordBox uses PrintCoord and then creates a box around it.
func (w *World) PrintCoordBox() (string, error) {
	inner, err0 := w.PrintCoord()
	if err0 != nil {
		return "", err0
	}
	boxed, err1 := makeBox(w.Lat, w.Lon, inner, w.Inactive)
	if err1 != nil {
		return "", err1
	}
	return boxed, nil
}

// PrintCoord calculates x and y coordinates within 2D string world map from latitude and longitude values and marks the location within the 2D string.
func (w *World) PrintCoord() (string, error) {
	col, err0 := lon2col(w.Lon)
	row, err1 := lat2row(w.Lat)
	if err0 != nil {
		return MAP, err0
	}
	if err1 != nil {
		return MAP, err1
	}
	markedMap, err2 := markMap(col, row, MAP, markerSpot, w.Inactive)
	return markedMap, err2
}

// ConcatenateHorizontally concatenates two multi-line strings horizontally
func ConcatenateHorizontally(str1, str2 string) string {
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
