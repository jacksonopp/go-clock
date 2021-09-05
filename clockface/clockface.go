package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

const (
	secondsInHalfClock float64 = 30
	secondsInClock     float64 = 2 * secondsInHalfClock
	minutesInHalfClock float64 = 30
	minutesInClock     float64 = 2 * minutesInHalfClock
	hoursInHalfClock   float64 = 6
	hoursInClock       float64 = 2 * hoursInHalfClock
)

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	return AngleToPoint(angle)
}

func MinuteHandPoint(t time.Time) Point {
	angle := MinutesInRadians(t)
	return AngleToPoint(angle)
}

func HourHandPoint(t time.Time) Point {
	angle := HoursInRadians(t)
	return AngleToPoint(angle)
}

func AngleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
