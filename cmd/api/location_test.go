package main

import (
	"errors"
	"testing"
)

func TestNewLocation(t *testing.T) {
	t.Run("test valid ipv4 address", func(t *testing.T) {
		ip := "71.38.107.126"
		loc, err := newLocation(ip)
		if err != nil {
			t.Fatalf("got unexpected error")
		}

		if !validLocation(loc) {
			t.Errorf("location is invalid")
		}
	})

	t.Run("test valid ipv6 address", func(t *testing.T) {
		ip := "2001:4860:4860::8888"
		loc, err := newLocation(ip)
		if err != nil {
			t.Fatalf("got unexpected error")
		}

		if !validLocation(loc) {
			t.Errorf("location is invalid")
		}

	})

	t.Run("test invalid address", func(t *testing.T) {
		ip := "5"
		_, err := newLocation(ip)
		if nil == err {
			t.Fatalf("expected error but didn't get one")
		}

		if !errors.Is(err, errInvalidIP) {
			t.Fatalf("expected %v, got %v", errInvalidIP, err)
		}
	})

	t.Run("test empty address", func(t *testing.T) {
		ip := ""
		loc, err := newLocation(ip)
		if err != nil {
			t.Fatalf("got unexpected error")
		}

		if !validLocation(loc) {
			t.Errorf("location is invalid")
		}
	})
}

func TestMidpoint(t *testing.T) {
	tests := []struct {
		loc1    *location
		loc2    *location
		wantLat int
		wantLon int
	}{
		{
			&location{0, 0},
			&location{90, 180},
			45, 0,
		},
		{
			&location{-90, -180},
			&location{90, 180},
			0, 180,
		},
	}

	for _, v := range tests {
		res := midpoint(v.loc1, v.loc2)
		gotLat := int(res.Latitude)
		gotLon := int(res.Longitude)

		if gotLat != v.wantLat || gotLon != v.wantLon {
			t.Errorf("incorrect midpoint for %v and %v",
				v.loc1, v.loc2)
		}
	}
}

func validLocation(loc *location) bool {
	if loc == nil {
		return false
	}

	if loc.Latitude < -90.0 || loc.Latitude > 90.0 {
		return false
	}

	if loc.Longitude < -180.0 || loc.Longitude > 180.0 {
		return false
	}

	return true
}
