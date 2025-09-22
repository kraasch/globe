
# geo

## Demo

Demo picture:

...

<!--
<p align="center">
<img src="./resources/example.png" width="300"/>
</p>
-->

## Inspiration

Example print out of some data.

```text
now local:    2025-09-22 20:45:53 +02:00
now utc:      2025-09-22 18:45:53.623185
```

See inspiration scripts

  - [./inspiration/moon](./inspiration/moon)
  - [./inspiration/sun](./inspiration/sun)

## Features

List of features

  - [ ] xxx

## Tasks

List of things to do

  - [ ] conversion function from location+time to timezone (default to system timezone).
  - [ ] visualizations.
    - [ ] visualization of lat+lon (user location).
    - [ ] visualization of moon position with longitude.
    - [ ] visualization of sun position with longitude.

List of things done

  - [X] find out where user is located (ask system).
  - [X] find web for lat+lon (ask web).
  - [X] pretty print earth.
  - [X] calculate today's hours of sunrise and sunset.
  - [X] calculate current phase of moon.

Ideas:

  - [ ] pull out TUI world map into its own package, eg `tuiworld`

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

## License

View the [license file](./LICENSE).

