package clock

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock

	secondHandLength = 90.0
	minuteHandLength = 80.0
	hourHandLength   = 50.0
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

//XML2struct: https://www.onlinetool.io/xmltogo/
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func secondHand(w io.Writer, t time.Time) {
	unitCirclePoint := secondHandPoint(t)
	point := makeHand(secondHandLength, unitCirclePoint)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, point.X, point.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	unitCirclePoint := minuteHandPoint(t)
	point := makeHand(minuteHandLength, unitCirclePoint)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:4px;"/>`, point.X, point.Y)
}

func hourHand(w io.Writer, t time.Time) {
	unitCirclePoint := hourHandPoint(t)
	point := makeHand(hourHandLength, unitCirclePoint)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, point.X, point.Y)
}

func makeHand(handLength float64, unitCirclePoint Point) Point {
	point := Point{
		X: clockCentreX + handLength*unitCirclePoint.X,
		Y: clockCentreY - handLength*unitCirclePoint.Y,
	}
	return point
}

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondsInRadians(t time.Time) float64 {
	val := float64(t.Second())
	return math.Pi / (secondsInHalfClock / val)
}

func minutesInRadians(t time.Time) float64 {
	val := float64(t.Minute()*minutesInClock + t.Second())
	return math.Pi / ((secondsInHalfClock * minutesInClock) / val)
}

func hoursInRadians(t time.Time) float64 {
	val := float64((t.Hour()%hoursInClock)*secondsInClock*minutesInClock + t.Minute()*minutesInClock + t.Second())
	return math.Pi / ((secondsInHalfClock * minutesInClock * hoursInClock) / val)
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{
		X: math.Sin(angle),
		Y: math.Cos(angle),
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(7, time.March, 13, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
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
