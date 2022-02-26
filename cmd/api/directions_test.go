package main

import (
	"errors"
	"os"
	"testing"
)

var testApp *application

func TestMain(m *testing.M) {
	cfg := config{
		mapsAPIkey: os.Getenv("MAPKEY"),
	}

	testApp = &application{
		cfg: cfg,
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestNewDirections(t *testing.T) {
	t.Run("test directions between valid locations", func(t *testing.T) {
		origin := &location{36.1147074, -115.1728497}
		dest := &location{35.67591606795329, -110.86160042036988}

		_, err := testApp.newDirections(origin, dest)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	})

	t.Run("test directions to invalid location", func(t *testing.T) {
		origin := &location{0, 0}
		dest := &location{95, 360}

		_, err := testApp.newDirections(origin, dest)
		if nil == err {
			t.Fatalf("expected an error but didn't get one")
		}

		if !errors.Is(err, errDirectionsNotFound) {
			t.Fatalf("expected %s, got %s", errDirectionsNotFound,
				err)
		}
	})
}

func TestGetAddress(t *testing.T) {
	t.Run("test get address of invalid location", func(t *testing.T) {
		loc := &location{100, -115.1728497}
		_, err := testApp.getAddress(loc)
		if nil == err {
			t.Fatalf("expected an error but didn't get one")
		}

		if !errors.Is(err, errDirectionsNotFound) {
			t.Fatalf("expected %s, got %s", errAddressNotFound,
				err)
		}
	})
}
