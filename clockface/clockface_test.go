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
			got := secondsInRadians(c.time)
			want := c.angle

			if want != got {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{-1, 0}},
		{simpleTime(0, 0, 45), Point{0, -1}},
	}

	for _, c := range cases {
		t.Run(helpers.TestName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			want := c.point

			if roughlyEqualPoints(got, want) {
				t.Fatalf("Got point %v, wanted point %v", got, want)
			}
		})
	}
}
