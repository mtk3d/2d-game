package main

import (
	"strconv"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Location struct {
	walls []Wall
}

func NewLocation(walls []Wall) Location {
	return Location{
		walls: walls,
	}
}

// location string - x,y;x,y;break;x,y;
func BuildLocation(location string) Location {
	points := strings.Split(location, ";")
	var (
		walls      []Wall
		wallPoints []pixel.Vec
	)
	for _, point := range points {
		if point == "break" {
			walls = append(walls, NewWall(wallPoints))
			wallPoints = nil
		} else {
			cords := strings.Split(point, ",")
			x, _ := strconv.ParseFloat(cords[0], 64)
			y, _ := strconv.ParseFloat(cords[1], 64)
			wallPoints = append(wallPoints, pixel.V(x, y))
		}
	}

	walls = append(walls, NewWall(wallPoints))

	return NewLocation(walls)
}

func (l *Location) Move(move pixel.Vec) {
	for i := range l.walls {
		l.walls[i].Move(move)
	}
}

func (l *Location) Draw(win *pixelgl.Window) {
	for _, wall := range l.walls {
		wall.Draw(win)
	}
}
