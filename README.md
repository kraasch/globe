
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

  - [ ] find out what time it is.
  - [ ] find out where user is located (ask system or web).
    - [ ] convert to lat+lon.
  - [ ] visualizations.
    - [ ] visualization of lat+lon (user location).
    - [ ] visualization of moon position with longitude.
    - [ ] visualization of sun position with longitude.
  - [ ] conversion function from location+time to timezone.
    - [ ] find this on github if possible.
  - [ ] on earth print, show current location.

List of things done

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

## License

View the [license file](./LICENSE).

