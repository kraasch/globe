
# geo

## Demo

Demo picture:

<p align="center">
  <img src="./resources/example.png" width="300"/>
</p>

## Inspiration

Example print out of some data.

```text
now local:    2025-09-22 20:45:53 +02:00
now utc:      2025-09-22 18:45:53.623185
sunrise time: 2025-09-23 06:33:25.927160
sunset time:  2025-09-22 18:32:28.773175
lunar phase:   ○ New Moon
moon position: 0.018
      0: "○ New Moon",
      1: "❩ Waxing Crescent",
      2: "◗ First Quarter",
      3: "◑ Waxing Gibbous",
      4: "● Full Moon",
      5: "◐ Waning Gibbous",
      6: "◖ Last Quarter",
      7: "❨ Waning Crescent"

```

## Features

List of features

  - [ ] pretty print earth.
    - [ ] show current location.
  - [ ] show current phase of moon.
  - [ ] calculate today's hours of sunrise and sunset.

## Tasks

List of things to do

  - [ ] xxx

List of ideas

  - [ ] xxx

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

  - https://github.com/nathan-osman/go-sunrise

## License

View the [license file](./LICENSE).

