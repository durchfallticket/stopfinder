// Copyright 2022 Hugo Melder. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package stopfinder

type StopArea []struct {
	Elevation         int    `json:"elevation"`
	ParentName        string `json:"parentName"`
	Coord             Coord  `json:"coord"`
	Omc               int    `json:"omc"`
	Level             int    `json:"level"`
	Nets              []Nets `json:"nets"`
	IsTransferStation bool   `json:"isTransferStation"`
	Name              string `json:"name"`
	ID                string `json:"id"`
	ParentID          string `json:"parentId"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Nets struct {
	Name  string `json:"name"`
	Zones []int  `json:"zones"`
}

type DataVersion struct {
	Format   int    `json:"format"`
	Name     string `json:"name"`
	Date     int    `json:"date"`
	Time     int    `json:"time"`
	MapName  string `json:"mapName"`
	MapUnit  int    `json:"mapUnit"`
	CheckSum int    `json:"checkSum"`
	Stops    int    `json:"stops"`
  /* TODO: LINK_AREA and STOP_AREA_BEACONS parsers
   * are not implemented.
   *
	 * Links    int    `json:"links"`
	 * Beacons  int    `json:"beacons"`
   */
}
