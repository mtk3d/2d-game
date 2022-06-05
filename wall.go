package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	wallThick = 10
	halfThick = wallThick / 2
)

type Wall struct {
	points []pixel.Vec
}

func NewWall(points []pixel.Vec) Wall {
	var linePoints []pixel.Vec

	for i, point := range points {
		if i == len(points)-1 {
			prevPoint := points[i-1]
			angle := prevPoint.Sub(point).Angle()
			linePoints = append(linePoints, pixel.V(-halfThick, -halfThick).Rotated(angle).Add(point))
			linePoints = append(linePoints, pixel.V(-halfThick, halfThick).Rotated(angle).Add(point))
		} else {
			nextPoint := points[i+1]
			angle := nextPoint.Sub(point).Angle()
			linePoints = append(linePoints, pixel.V(-halfThick, halfThick).Rotated(angle).Add(point))
		}
	}

	for i := range points {
		ri := len(points) - i - 1
		point := points[ri]

		if ri == 0 {
			prevPoint := points[ri+1]
			angle := prevPoint.Sub(point).Angle()
			linePoints = append(linePoints, pixel.V(-halfThick, -halfThick).Rotated(angle).Add(point))
			linePoints = append(linePoints, pixel.V(-halfThick, halfThick).Rotated(angle).Add(point))
		} else {
			nextPoint := points[ri-1]
			angle := nextPoint.Sub(point).Angle()
			linePoints = append(linePoints, pixel.V(halfThick, halfThick).Rotated(angle).Add(point))
		}
	}

	return Wall{
		points: linePoints,
	}
}

func (w *Wall) Move(move pixel.Vec) {
	for i, point := range w.points {
		w.points[i] = point.Add(move)
	}
}

func (w *Wall) Draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Gray
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(w.points...)
	imd.Line(1)
	imd.Draw(win)
}
