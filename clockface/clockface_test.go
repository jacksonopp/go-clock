package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/jacksonopp/go-clock/helpers"
)

func roughlyEqualFloats(a, b float64) bool {
	const equalityThresh = 1e-7
	return math.Abs(a-b) < equalityThresh
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloats(a.X, b.X) && roughlyEqualFloats(a.Y, b.Y)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{helpers.SimpleTime(0, 0, 30), math.Pi},
		{helpers.SimpleTime(0, 0, 0), 0},
		{helpers.SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{helpers.SimpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)
			want := c.angle

			if !roughlyEqualFloats(want, got) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{helpers.SimpleTime(0, 30, 0), math.Pi},
		{helpers.SimpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := MinutesInRadians(c.time)
			if !roughlyEqualFloats(got, c.angle) {
				t.Fatalf("Wanted %v radians, got %v", c.angle, got)
			}
		})
	}
}
func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{helpers.SimpleTime(6, 0, 0), math.Pi},
		{helpers.SimpleTime(0, 0, 0), 0},
		{helpers.SimpleTime(21, 0, 0), math.Pi * 1.5},
		{helpers.SimpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := HoursInRadians(c.time)
			if !roughlyEqualFloats(got, c.angle) {
				t.Fatalf("Wanted %v radians, got %v", c.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{helpers.SimpleTime(0, 0, 30), Point{-1, 0}},
		{helpers.SimpleTime(0, 0, 45), Point{0, -1}},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			want := c.point

			if roughlyEqualPoints(got, want) {
				t.Fatalf("Got point %v, wanted point %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{helpers.SimpleTime(0, 30, 0), Point{0, -1}},
		{helpers.SimpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := MinuteHandPoint(c.time)
			if !roughlyEqualPoints(got, c.point) {
				t.Fatalf("Wanted %v Point but got %v", c.point, got)
			}
		})
	}
}
func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{helpers.SimpleTime(6, 0, 0), Point{0, -1}},
		{helpers.SimpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := HourHandPoint(c.time)
			if !roughlyEqualPoints(got, c.point) {
				t.Fatalf("Wanted %v Point but got %v", c.point, got)
			}
		})
	}
}
