// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type App struct {
	uint8
}

func NewApp(arg []string) *App {
	return &App{}
}

func (g *App) Update() error {

	g.uint8++

	return nil
}

func (g *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: uint8(g.uint8), G: uint8(g.uint8), B: 23, A: uint8(g.uint8)})
}

func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 250, 90
}
