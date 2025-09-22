// Package geoshow is used to nicely print information calculated or retreiveed from geocalc and geoauto.
package geoshow

func PrintEarth() string {
	NL := "\n"
	return "┌────────────────────────┐" + NL +
		"│1 9876-4321 1234+6789 12│" + NL +
		"├───────────▼────────────┤" + NL +
		"│    _,--._  _._.--.--.._│" + NL +
		"▶=.--'=_',-,▣`;_      .,'◀" + NL +
		"│,-.  _.)  (``-;_   .'   │" + NL +
		"│   '-:_    `) ) v''=.   │" + NL +
		"│     ) )    ()'    ='   │" + NL +
		"│     |/            (_) =│" + NL +
		"├───────────▲────────────┤" + NL +
		"│   ☼            ●       │" + NL +
		"└────────────────────────┘"
}
