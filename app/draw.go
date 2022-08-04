// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: uint8(g.uint8), G: uint8(g.uint8), B: 23, A: uint8(g.uint8)})

	UiLayout(screen)
}
