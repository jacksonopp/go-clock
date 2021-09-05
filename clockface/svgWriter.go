package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength float64 = 90
	minuteHandLength float64 = 80
	hourHandLength   float64 = 50
	clockCenter      float64 = 150
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	HourHand(w, t)
	io.WriteString(w, svgEnd)
}

func MakeHand(w io.Writer, p *Point, length float64, color string) {
	*p = Point{p.X * length, p.Y * length}
	*p = Point{p.X, -p.Y}
	*p = Point{p.X + clockCenter, p.Y + clockCenter}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`, p.X, p.Y, color)
}

func SecondHand(w io.Writer, t time.Time) {
	p := SecondHandPoint(t)
	MakeHand(w, &p, secondHandLength, "#f00")
}

func MinuteHand(w io.Writer, t time.Time) {
	p := MinuteHandPoint(t)
	MakeHand(w, &p, minuteHandLength, "#000")
}

func HourHand(w io.Writer, t time.Time) {
	p := HourHandPoint(t)
	MakeHand(w, &p, hourHandLength, "#000")
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
