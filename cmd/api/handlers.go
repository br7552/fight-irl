package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/br7552/router"
)

func (app *application) yourAddrInfoHandler(w http.ResponseWriter,
	r *http.Request) {

	yourIP := getIP(r)
	if strings.HasPrefix(yourIP, "127.0.0.1") {
		yourIP = ""
	}

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

func (app *application) theirAddrInfoHandler(w http.ResponseWriter,
	r *http.Request) {

	theirIP := router.Param(r, "ip")
	theirLoc, err := newLocation(theirIP)
	if err != nil {
		if errors.Is(err, errInvalidIP) {
			app.badRequestResponse(w, r, errInvalidIP)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
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
		IP      string    `json:"ip"`
		Coords  *location `json:"coordinates"`
		Address string    `json:"address"`
	}{
		theirIP,
		theirLoc,
		theirAddress,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"their_address_information": data,
	}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) meetingHandler(w http.ResponseWriter,
	r *http.Request) {

	yourIP := getIP(r)
	if strings.HasPrefix(yourIP, "127.0.0.1") {
		yourIP = ""
	}

	yourLoc, err := newLocation(yourIP)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	theirIP := router.Param(r, "ip")
	theirLoc, err := newLocation(theirIP)
	if err != nil {
		if errors.Is(err, errInvalidIP) {
			app.badRequestResponse(w, r, errInvalidIP)
		} else {
			app.serverErrorResponse(w, r, err)
		}
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
		MeetingLocation *location   `json:"meeting_location"`
		YourAddress     string      `json:"your_start_address"`
		TheirAddress    string      `json:"their_start_address"`
		YourDirections  *directions `json:"your_directions,omitempty"`
		TheirDirections *directions `json:"their_directions,omitempty"`
	}{
		dest,
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

func getIP(r *http.Request) string {
	forwarded := strings.SplitN(r.Header.Get("X-FORWARDED-FOR"), ",", 2)
	if len(forwarded) > 0 {
		return forwarded[0]
	}

	return r.RemoteAddr
}
