
# geo

An uncluttered, minimal and clean display of commonly used every-day astronomical data about the moon, sun, earth and the computer's location.

## Demo

Demo picture:

<p align="center">
<img src="./resources/example.png" width="600"/>
</p>

## Features

List of features

  - [ ] xxx

## Tasks

List of things to do

  - [ ] make sure every data in the display is read in automatically.
  - [ ] data in display updates on interval.
  - [ ] provide an option for an update interval of all displayed data.

List of things done

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

## Usage

Use the program:

```bash
make build
./build/geo --help
```

Use the package:

```go
import (
  "github.com/kraasch/geo"
)

geo.DoSomething("Hello")
```

## Inspiration

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

## Misc Info

Other useful packages:

  - for sunrise and sunset: https://github.com/nathan-osman/go-sunrise
  - for moon phases: https://github.com/janczer/goMoonPhase
  - for lon+lat to timezone conversion: https://github.com/evanoberholster/timezoneLookup
  - more symbols: https://en.wikipedia.org/wiki/Geometric_Shapes_(Unicode_block)

## License

View the [license file](./LICENSE).

