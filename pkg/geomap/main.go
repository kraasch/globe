// Package geomap is is used to show an ASCII world map on the CLI.
package geomap

import (
	"fmt"
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

func NewWorld() World {
	return World{"default"}
}

func (w *World) PrintCoord(lat, lon float64) (string, error) {
	return MAP, nil
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
