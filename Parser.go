package stopfinder

import (
	"strings"
  "math"
)

func ParseBytesFormat5(bArr *[]byte) []Info {
	br := Br{Data: *bArr, Pointer: 0}
	var arrList []Info

	for HasNext(&br) {
		count := GetInt(&br)
		for i2 := 0; i2 < count; i2++ {
			parentId := GetVariableLengthString(&br)
			parentName := GetVariableLengthString(&br)
			omc := GetInt(&br)
			coord := NewCoordinate(float64(GetInt(&br))/100.0, float64(GetInt(&br))/100.0)
			netLength := GetByte(&br)
			nets := make([]Net, int(math.Max(0, float64(netLength - 1))))
			for netIndex := 1; netIndex <= netLength; netIndex++ {
				netName := strings.ToLower(GetString(&br, 3))
				zoneLength := GetByte(&br)
				zones := make([]int, zoneLength-1)
				for zoneIndex := 1; zoneIndex <= zoneLength; zoneIndex++ {
					zones = append(zones, GetInt(&br))
				}
				nets = append(nets, Net{Name: netName, Zones: zones})
			}
			length := int(GetShort(&br))
			for index := 1; index <= length; index++ {
				id := GetVariableLengthString(&br)
				name := GetVariableLengthString(&br)
				level := GetShort(&br)
				elevation := GetShort(&br)
				isTransferStation := GetByte(&br) == 1
				length2 := GetByte(&br)
				for index2 := 1; index2 <= int(length2); index2++ {
					GetInt(&br)
					GetInt(&br)
				}
				arrList = append(arrList, Info{
					Coord:             coord,
					Elevation:         elevation,
					Id:                id,
					IsTransferStation: isTransferStation,
					Level:             level,
					Name:              name,
					Nets:              nets,
					Omc:               omc,
					ParentId:          parentId,
					ParentName:        parentName,
					NormalizedName:    normalizeStr(parentName)})
			}
		}
	}
	return arrList
}
