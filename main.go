package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Getting Started with Ebiten")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(100)

	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}