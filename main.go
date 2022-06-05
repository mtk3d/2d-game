package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	width     = 800
	height    = 600
	moveSpeed = 7
)

var (
	center          = pixel.V(width/2, height/2)
	backgroundColor = colornames.Black
	character       Character
)

func loadLocation() Location {
	content, err := ioutil.ReadFile("location.txt")

	if err != nil {
		log.Fatal(err)
	}

	return BuildLocation(string(content))
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "2D Game Engine",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(30 * time.Millisecond)

	character = NewCharacter(center, 2)

	location := loadLocation()

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyA) {
			location.Move(pixel.V(1*moveSpeed, 0))
		}
		if win.Pressed(pixelgl.KeyD) {
			location.Move(pixel.V(-1*moveSpeed, 0))
		}
		if win.Pressed(pixelgl.KeyS) {
			location.Move(pixel.V(0, 1*moveSpeed))
		}
		if win.Pressed(pixelgl.KeyW) {
			location.Move(pixel.V(0, -1*moveSpeed))
		}

		character.RotateTo(win.MousePosition())

		select {
		case <-tick:
			win.Clear(backgroundColor)
			character.Draw(win)
			location.Draw(win)
			win.Update()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
