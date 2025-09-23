// Package geomap is is used to show an ASCII world map on the CLI.
package geomap

type World struct {
	name string // TODO: delete later.
}

func NewWorld() World {
	return World{"default"}
}

func (w *World) PrintBlank() string {
	NL := "\n"
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
	NL := "\n"
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
