package main

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

var ipWhoisURL = "http://ipwhois.app/json/"

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func newLocation(ip string) (*location, error) {
	url := ipWhoisURL + ip
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loc location
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return nil, err
	}

	return &loc, nil
}

func midpoint(loc1, loc2 *location) *location {
	lat1 := loc1.Latitude * (math.Pi / 180)
	lon1 := loc1.Longitude * (math.Pi / 180)

	lat2 := loc2.Latitude * (math.Pi / 180)
	lon2 := loc2.Longitude * (math.Pi / 180)

	x1 := math.Cos(lat1) * math.Cos(lon1)
	y1 := math.Cos(lat1) * math.Sin(lon1)
	z1 := math.Sin(lat1)

	x2 := math.Cos(lat2) * math.Cos(lon2)
	y2 := math.Cos(lat2) * math.Sin(lon2)
	z2 := math.Sin(lat2)

	x := (x1 + x2) / 2
	y := (y1 + y2) / 2
	z := (z1 + z2) / 2

	lon := math.Atan2(y, x)
	hyp := math.Sqrt(x*x + y*y)
	lat := math.Atan2(z, hyp)

	return &location{
		Latitude:  lat * (180 / math.Pi),
		Longitude: lon * (180 / math.Pi),
	}
}
