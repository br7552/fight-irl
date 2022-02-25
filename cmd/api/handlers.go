package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/br7552/router"
)

func (app *application) addrInfoHandler(w http.ResponseWriter,
	r *http.Request) {

	//yourIP := r.RemoteAddr
	yourIP := "172.58.76.209"

	yourLoc, err := newLocation(yourIP)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	yourAddress, err := app.getAddress(yourLoc)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			yourAddress = "unknown"
		} else {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	data := struct {
		IP      string    `json:"ip"`
		Coords  *location `json:"coordinates"`
		Address string    `json:"address"`
	}{
		yourIP,
		yourLoc,
		yourAddress,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"your_address_information": data,
	}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) meetingHandler(w http.ResponseWriter,
	r *http.Request) {

	//yourIP := r.RemoteAddr
	yourIP := "172.58.76.209"
	theirIP := router.Param(r, "ip")
	// validate ip
	// if ip format invalid, return bad request response

	yourLoc, err := newLocation(yourIP)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	theirLoc, err := newLocation(theirIP)
	if err != nil {
		log.Println(err)
		return
	}

	dest := midpoint(yourLoc, theirLoc)

	yourDirections, err := app.newDirections(yourLoc, dest)
	if err != nil {
		if !errors.Is(err, errDirectionsNotFound) {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	theirDirections, err := app.newDirections(theirLoc, dest)
	if err != nil {
		if !errors.Is(err, errDirectionsNotFound) {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	destAddress, err := app.getAddress(dest)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			destAddress = "unknown"
		} else {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	yourAddress, err := app.getAddress(yourLoc)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			yourAddress = "unknown"
		} else {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	theirAddress, err := app.getAddress(theirLoc)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			theirAddress = "unknown"
		} else {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	data := struct {
		MeetingLocation string      `json:"meeting_location"`
		YourLocation    string      `json:"your_location"`
		TheirLocation   string      `json:"their_location"`
		YourDirections  *directions `json:"your_directions,omitempty"`
		TheirDirections *directions `json:"their_directions,omitempty"`
	}{
		destAddress,
		yourAddress,
		theirAddress,
		yourDirections,
		theirDirections,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"meeting": data,
	}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
