package stopfinder

import (
  "os"
  "math"
)

const (
  /* The STOP_AREA.bin file has an offset
   * before real data appears
   */
  BinaryOffset = 162193
)

type Stopfinder interface {
  Find(name string) *[]Info
}

type DefaultStopfinder struct {
  Set      *FuzzySet
  Stops    []Info
}

func createNameSlice(stops *[]Info) []string {
	parentNames := make([]string, len(*stops))
	for _, obj := range *stops {
		parentNames = append(parentNames, obj.NormalizedName)
	}

  return parentNames
}

func NewStopfinder(blob string) (*DefaultStopfinder, error) {
  var stopfinder DefaultStopfinder

  data, err := os.ReadFile(blob)
  if err != nil {
    return nil, err
  } 
  // Apply the offset
  data = data[BinaryOffset:]

  // Parse the blob
  stopfinder.Stops = ParseBytesFormat5(&data)

  nameSlice := createNameSlice(&stopfinder.Stops)
  stopfinder.Set = New(nameSlice, false, 2, 3, 0.6)

  return &stopfinder, nil
}

func (d *DefaultStopfinder) Find(name string) *[]Info {
  var stops []Info

  fuzzyResults := d.Set.Get(name)
	
  for _, match := range fuzzyResults[:int(math.Min(10, float64(len(fuzzyResults))))] {
		for _, obj := range d.Stops {
			if match.Word == obj.NormalizedName {
        stops = append(stops, obj)
			}
		}
	}

  return &stops
}
