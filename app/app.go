// Created by EldersJavas(EldersJavas&gmail.com)

package app

type App struct {
	uint8
}

func NewApp(arg []string) *App {
	return &App{}
}

func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 250, 90
}
