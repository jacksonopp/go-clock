package clockface_test

import (
	"testing"
	"time"

	"github.com/jacksonopp/go-clock/clockface"
)

// func TestClockface(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 - 90}
// 	got := clockface.SecondHand(tm)

// 	if want != got {
// 		t.Errorf("got %v, want %v", want, got)
// 	}
// }

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
