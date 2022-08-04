// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/EldersJavas/EbiSplit/app"
	"github.com/EldersJavas/EbiSplit/config"
	"github.com/beevik/ntp"
	"github.com/hajimehoshi/ebiten/v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var AppConfigs config.Config

func init() {
	co, err := ioutil.ReadFile("config.json")
	AppConfigs, _ = config.UnmarshalConfig(co)
	if err != nil {
		log.Println(err)
	}
	ebiten.SetWindowSize(AppConfigs.WindowSize.Width, AppConfigs.WindowSize.Height)
	ebiten.SetMaxTPS(10)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowTitle("EbiSplit")

	//ebiten.SetWindowSizeLimits(250, 90, 500, 180)
	ebiten.SetWindowDecorated(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			RunApp()
		}
	}()

	options := ntp.QueryOptions{Timeout: 3 * time.Second}
	response, err := ntp.QueryWithOptions("0.pool.ntp.org", options)
	if err != nil {
		log.Println(err)
	}
	if response != nil {
		time1 := time.Now().Add(response.RTT)
		log.Println(time1)
	}
	RunApp()

}
func RunApp() {

	err := ebiten.RunGame(app.NewApp(os.Args))
	if err != nil {
		log.Fatal(err)
	}
}
