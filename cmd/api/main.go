package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type config struct {
	mapsAPIkey string
}

type application struct {
	cfg config
}

func main() {
	cfg := config{
		mapsAPIkey: os.Getenv("MAPKEY"),
	}

	app := &application{
		cfg: cfg,
	}

	lambda.Start(app.meetingHandler)
}
