package main

import (
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
		log.Println(err)
		return
	}

	yourAddress := app.getAddress(yourLoc)

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
		log.Println(err)
	}
}

func (app *application) meetingHandler(w http.ResponseWriter,
	r *http.Request) {

	//yourIP := r.RemoteAddr
	yourIP := "172.58.76.209"
	theirIP := router.Param(r, "ip")

	yourLoc, err := newLocation(yourIP)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}

	theirDirections, err := app.newDirections(theirLoc, dest)
	if err != nil {
		log.Println(err)
		return
	}

	data := struct {
		MeetingLocation string      `json:"meeting_location"`
		YourLocation    string      `json:"your_location"`
		TheirLocation   string      `json:"their_location"`
		YourDirections  *directions `json:"your_directions"`
		TheirDirections *directions `json:"their_directions"`
	}{
		app.getAddress(dest),
		app.getAddress(yourLoc),
		app.getAddress(theirLoc),
		yourDirections,
		theirDirections,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"meeting": data,
	}, nil)
	if err != nil {
		log.Println(err)
	}
}
