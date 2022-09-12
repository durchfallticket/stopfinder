package stopfinder

import (
  "testing"
  "fmt"
)

func TestStopfinder(t *testing.T) {
  sfinder, err := NewStopfinder("STOP_AREA.bin")
  if err != nil {
    t.Fatal(err)
  }

  results := sfinder.Find("Düsseldorf Hbf")
  fmt.Printf("%+v\n", results)
}
