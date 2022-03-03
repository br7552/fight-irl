package main

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func (app *application) meetingHandler(ctx context.Context,
	r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,
	error) {

	yourIP := getIP(r)
	if strings.HasPrefix(yourIP, "127.0.0.1") {
		yourIP = ""
	}

	yourLoc, err := newLocation(yourIP)
	if err != nil {
		return serverErrorResponse(err)
	}

	theirIP := r.PathParameters["ip"]
	theirLoc, err := newLocation(theirIP)
	if err != nil {
		if errors.Is(err, errInvalidIP) {
			return badRequestResponse(errInvalidIP)
		} else {
			return serverErrorResponse(err)
		}
	}

	dest := midpoint(yourLoc, theirLoc)

	yourDirections, err := app.newDirections(yourLoc, dest)
	if err != nil {
		if !errors.Is(err, errDirectionsNotFound) {
			return serverErrorResponse(err)
		}
	}

	theirDirections, err := app.newDirections(theirLoc, dest)
	if err != nil {
		if !errors.Is(err, errDirectionsNotFound) {
			return serverErrorResponse(err)
		}
	}

	yourAddress, err := app.getAddress(yourLoc)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			yourAddress = "unknown"
		} else {
			return serverErrorResponse(err)
		}
	}

	theirAddress, err := app.getAddress(theirLoc)
	if err != nil {
		if errors.Is(err, errAddressNotFound) {
			theirAddress = "unknown"
		} else {
			return serverErrorResponse(err)
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

	return writeJSON(http.StatusOK, envelope{
		"meeting": data,
	})
}

func getIP(r events.APIGatewayProxyRequest) string {
	forwarded := strings.SplitN(r.Headers["X-Forwarded-For"], ",", 2)
	if len(forwarded) > 0 {
		return forwarded[0]
	}

	return r.RequestContext.Identity.SourceIP
}
