// Copyright 2022 Hugo Melder. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package stopfinder

/* A public transport stop based on the IFOPT standard.
 */
type Stop struct {
	// Elevation of the public transport stop
	Elevation int `json:"elevation"`
	// Coordinates of the public transport stop
	Coord Coord `json:"coord"`
	Omc   int   `json:"omc"`
	Level int   `json:"level"`
	// The network that operates the stop
	Nets              []Nets `json:"nets"`
	IsTransferStation bool   `json:"isTransferStation"`
	// The name of the stop
	Name string `json:"name"`
	// The name of the parent transport stop (when nested)
	ParentName string `json:"parentName"`
	// The global identifier of the stop (e.g. 'de:05111:18235')
	ID string `json:"id"`
	// The global identifier of the parent transport stop (when nested)
	ParentID string `json:"parentId"`
}

type Coord struct {
	// Latitude
	Lat float64 `json:"lat"`
	// Longitude
	Lon float64 `json:"lon"`
}

type Nets struct {
	// The name of the public transport network (e.g. 'vrs')
	Name string `json:"name"`
	// The zone(s) of the public transport network
	Zones []int `json:"zones"`
}

/* An object that describes the dataset.
 */
type DataVersion struct {
	// The internal format of the binary blob
	Format int `json:"format"`
	// Name of the dataset
	Name string `json:"name"`
	// Date when the dataset was created (e.g. '20220405')
	Date int `json:"date"`
	// Time when the dataset was created (e.g. '113534')
	Time int `json:"time"`
	// Name of the map
	MapName string `json:"mapName"`
	MapUnit int    `json:"mapUnit"`
	// Number of stops in STOP_AREA.bin
	Stops int `json:"stops"`
	/* TODO: LINK_AREA.bin and STOP_AREA_BEACONS.bin parsers
	   * are not implemented.
	   *
		 * Links    int    `json:"links"`
		 * Beacons  int    `json:"beacons"`
	*/
}
