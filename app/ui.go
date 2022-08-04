// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

/*
UI布局

--------------------------------
|   Time                        |
|                               |
|  GameName               icon  |
|  Best Record : xxx            |
-----------------------------------

*/
var (
	TimeFontImg     = ebiten.NewImage(250, 60)
	GameNameFontImg = ebiten.NewImage(250, 60)
	BestRFontImg    = ebiten.NewImage(250, 60)
	IconImg         = ebiten.NewImage(250, 60)
	FillOnImg       = ebiten.NewImage(250, 60)
	FillUnderImg    = ebiten.NewImage(250, 60)
)

func UiLayout(screen *ebiten.Image) {
	sw, sh := screen.Size()
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(sw/2-rand.Intn(300)+rand.Intn(300)), float64(sh/2-rand.Intn(300)+rand.Intn(300)))
	screen.DrawImage(emoji.Image("🎮"), op2)
}

func CreateDrawOp() (opt ebiten.DrawImageOptions) {
	opt = ebiten.DrawImageOptions{}

	return
}

/*file, err := ebitenutil.OpenFile("config.json")
if err != nil {
return
}
defer func(file ebitenutil.ReadSeekCloser) {
	err := file.Close()
	if err != nil {
		return
	}
}(file)
all, err := io.ReadAll(file)
if err != nil {
return
}
fmt.Println(string(all))

if err != nil {
return
}*/
