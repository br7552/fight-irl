package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var mapsURL = "https://maps.googleapis.com/maps/api/directions/json?"
var geocodeURL = "https://maps.googleapis.com/maps/api/geocode/json?"

type directions struct {
	Steps []step `json:"steps"`
}

type step struct {
	Text     string `json:"text"`
	Distance string `json:"distance"`
}

func (app *application) newDirections(start, dest *location) (*directions,
	error) {

	origin := fmt.Sprintf("origin=%f,%f",
		start.Latitude, start.Longitude)
	destination := fmt.Sprintf("&destination=%f,%f",
		dest.Latitude, dest.Longitude)
	key := fmt.Sprintf("&key=%s", app.cfg.mapsAPIkey)
	url := mapsURL + origin + destination + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var input struct {
		Status string `json:"status"`
		Routes []struct {
			Legs []struct {
				Steps []struct {
					Instructions string `json:"html_instructions"`
					Distance     struct {
						Text string `json:"text"`
					} `json:"distance"`
				} `json:"steps"`
			} `json:"legs"`
		} `json:"routes"`
	}

	/*
		if input.Status != "OK" {
			return nil, errors.New(input.Status)
		}
	*/
	err = json.Unmarshal(body, &input)
	if err != nil {
		return nil, err
	}

	var d directions

	legs := input.Routes[0].Legs
	for _, v := range legs {
		for _, u := range v.Steps {
			t := step{
				Text:     u.Instructions,
				Distance: u.Distance.Text,
			}
			d.Steps = append(d.Steps, t)
		}
	}

	return &d, nil
}

func (app *application) getAddress(loc *location) string {
	latlng := fmt.Sprintf("latlng=%f,%f",
		loc.Latitude, loc.Longitude)
	key := fmt.Sprintf("&key=%s", app.cfg.mapsAPIkey)
	url := geocodeURL + latlng + key

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var input struct {
		Results []struct {
			Address string `json:"formatted_address"`
		} `json:"results"`
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Println(err)
		return ""
	}

	return input.Results[0].Address
}
