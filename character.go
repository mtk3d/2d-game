package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Character struct {
	position        pixel.Vec
	angle           float64
	collisionPoints []pixel.Vec
}

func NewCharacter(position pixel.Vec, angle float64) Character {
	return Character{
		position: position,
		angle:    angle,
	}
}

func (c *Character) RotateTo(position pixel.Vec) {
	c.angle = position.Sub(c.position).Angle()
}

func (c *Character) Move(move pixel.Vec) {
	c.position = c.position.Add(move)
}

func (c *Character) Draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Limegreen
	imd.Push(c.position)
	imd.Circle(30, 1)
	imd.Push(c.position)
	imd.Push(pixel.V(50, 0).Rotated(c.angle).Add(c.position))
	imd.Line(1)
	imd.Draw(win)
}
