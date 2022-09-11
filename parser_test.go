package stopfinder

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	data, err := os.ReadFile("STOP_AREA.bin")
	if err != nil {
		fmt.Print(err)
	}
	data = data[162193:]
	result := ParseBytesFormat5(&data)
	fmt.Println("Parsed.")
	parentNames := make([]string, len(result))
	for _, obj := range result {
		parentNames = append(parentNames, obj.NormalizedName)
	}

	fmt.Println("Mapped.")

	fuzzy := New(parentNames, false, 2, 3, 0.6)

	fmt.Println("FuzzySet Created.")

	fuzzyResults := fuzzy.Get("Aachen Bushof")

	for _, match := range fuzzyResults[:int(math.Min(10, float64(len(fuzzyResults))))] {
		for _, obj := range result {
			if match.Word == obj.NormalizedName {
				fmt.Printf("%+v %f\n", obj, match.Score)
			}
		}
	}
}
