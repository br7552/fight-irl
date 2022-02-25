package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/br7552/router"
)

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
		YourDirections  *directions `json:"your_directions"`
		TheirDirections *directions `json:"their_directions"`
	}{
		app.getAddress(dest),
		yourDirections,
		theirDirections,
	}

	js, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
