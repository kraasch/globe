
package geoshow

import (
  "fmt"
)

const (
  value = "Toast: "
)

func Toast(in string) string {
  return fmt.Sprintf("%#v", value + in)
}

