
# geo

An uncluttered, minimal and clean display of commonly used every-day astronomical data about the moon, sun, earth and the computer's location.

## Demo and Basic Info

Demo picture:

<p align="center">
<img src="./resources/example.png" width="600"/>
</p>

User interaction on the TUI:

  - `u` for updating the data.
  - `q` for quitting the program.

## Features

List of features

  - [X] show commonly needed astronomical data at one glance.
  - [X] local calculations (except location of computer which uses web request).

## Usage Limits

This is an early version, the web requests are not buffered yet.
Executing the program or the tests too often will probably lead to being denied further requests.
See the section on usage limits on [ip-api.com/legal](https://ip-api.com/docs/legal) .

## Tasks

List of things to do

  - [ ] fix: buffer web retrieved lat+lon data as a text file somewhere.
  - [ ] fix: sunrise and sunset time is not accurate (1 or 2 hours off), match the local time in the timezone.
  - [ ] fix: sun longitude not calculated yet.
  - [ ] data in display updates on interval.
  - [ ] provide an option for an update interval of all displayed data.
  - [ ] refine tests and explicitly test data of
    - [ ] `▣ lat+lon: 53.48, 10.22`
    - [ ] `▣ zone:    Europe/Berlin`
    - [ ] `▣ offset:  UTC+2 (CEST)`
    - [ ] `▣ time:    12:52 h`
    - [ ] `🜃 utc:     10:52 h`
    - [ ] `☼ rise:    01:10 h`
    - [ ] `☼ set:     13:17 h`
    - [ ] `☼ lat+lon.`
    - [ ] `● phase:   Waning Gibbous (◐)`
    - [ ] `● age:     19.76 days`
    - [ ] `● dist.:   371578 km`
    - [ ] `● illum.:  74%`
    - [ ] `● new in:  10.6 days`
    - [ ] `● full in: 25.2 days`
    - [ ] `● new on:  2001-01-11, Thu`
    - [ ] `● full on: 2001-01-26, Fri`
    - [ ] `● lat+lon.`

List of things done

  - [X] make sure every data in the display is read in automatically.
  - [X] add time zone detection.
  - [X] find out user's time zone code, eg 'GMT'.
  - [X] find out user's utc shift for his time zone.
  - [X] find existing go package for timezone conversion (and maybe detection).
  - [X] make a unified, somewhat pretty display for all data.
  - [X] in geocalc: convert local time of user to UTC.
  - [X] in geocalc: conversion function from location+time to timezone (default to system timezone).
  - [X] visualizations.
    - [X] visualization of lat+lon (eg user location).
    - [X] visualization of moon position with longitude.
    - [X] visualization of sun position with longitude.
  - [X] pull out TUI world map into its own package (within the geo project), eg name the package `geomap`.
  - [X] find out where user is located (ask system).
  - [X] find web for lat+lon (ask web).
  - [X] pretty print earth.
  - [X] calculate today's hours of sunrise and sunset.
  - [X] calculate current phase of moon.

## Installation

Install the program:

```bash
go install github.com/kraasch/geo@latest
```

Get the package:

```bash
go get github.com/kraasch/geo
```

## In-Code Usage

Use the program:

```bash
make build
./build/geo --help
```

Use the package:

```go
import (
  geo "github.com/kraasch/geo"
)
var geoData geo.GeoData
geoData.PrintDataHorizontally()
```

## Future Inspiration

Some inspiring images for future development of this project:

  - [./inspiration/demo_0.png](./inspiration/demo_0.png)
  - [./inspiration/demo_1.png](./inspiration/demo_1.png)

## Feedback

I can be reached via [alex@kraasch.eu](mailto:alex@kraasch.eu).

## Contributing

Feel free to help me.

## Acknowledgments

Uses the following software:

  - see [go.mod](./go.mod) and [go.sum](./go.sum).

Made by the following people:

  - see Github info.

## Mini Astronomical Primer

**Geographic Coordinate System**
Data: Longitude and latitude.
Use: Stable reference on earth, navigating earth's surface, i.e. use of GPS or Google maps.

**Ecliptic Coordinate System**
Data: Ecliptic longitude and ecliptic latitude.
Use: Stable reference frame for sun and its planets, same for all observers on earth.

**Horizontal Coordinate System**
Data: Altitude (angle of an observed object), Azimuth (direction of the object along the horizon).
Use: References observed objects in the night sky, local to observer, changes during night with earth's rotation.

**Equatorial Coordinate System**
Data: Declination (how far north or south a celestial object is from the celestial equator), Right Ascension (hours, minutes and seconds along the celestial equator from a reference point called vernal equinox).
Use: Reference system fixed with respect to distant stars, slowly changing over long periods of time.

## Misc Info

Used packages:

  - for sunrise and sunset: https://github.com/nathan-osman/go-sunrise
  - for moon phases: https://github.com/janczer/goMoonPhase
  - for lon+lat to timezone conversion: https://github.com/evanoberholster/timezoneLookup
  - for timzeones: https://github.com/ringsaturn/tzf.

Other useful packages:

  - all kinds of astronomical algorithms: https://github.com/observerly
    - might contain sun position algorithm: https://github.com/observerly/sidera
  - for sun and moon position, dependent on the observer's location on earth: https://github.com/hablullah/go-sampa
    - NOTE: this package can calculate location-dependent sunrise and sunset.
  - can provide the sun's position (latitude and longitude): https://github.com/sj14/astral

Other useful resources:

  - more symbols: https://en.wikipedia.org/wiki/Geometric_Shapes_(Unicode_block)
  - good info on calculating sun's position: https://observablehq.com/@danleesmith/meeus-solar-position-calculations
  - info and implementation of some known algorithms: https://pkg.go.dev/github.com/soniakeys/meeus/v3
  - gist with usage example of soniakeys/meeus: https://gist.github.com/soniakeys/b066347d58a59ac6f3b4
  - info about different coordinate systems: https://pkg.go.dev/github.com/observerly/sidera@v0.7.0/pkg/coordinates
  - info on side real time: https://astro.dur.ac.uk/~ams/users/lst.html
  - calculator for sub-solar point from NASA: https://wgc.jpl.nasa.gov:8443/webgeocalc/#SubSolarPoint
  - web tool for visualization of sun and earth related data: https://www.sunearthtools.com/dp/tools/pos_sun.php?lang=en

## License

View the [license file](./LICENSE).

