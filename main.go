// Created by EldersJavas(EldersJavas&gmail.com)

package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	ebiten.SetMaxTPS(30)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowTitle("EbiSplit")
	ebiten.SetWindowDecorated(true)
}
