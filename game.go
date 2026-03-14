package main

import (
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
)

// embed asset
//go:embed asset
var assetFS embed.FS

const (
	screenW, screenH = 540, 304
	playerW, playerH = 16, 16
)

func loadImage(assetPath string) *ebiten.Image {
	f, err := assetFS.Open(assetPath)
	if err != nil {
		log.Panic(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type game struct {
	// background color
	bgColor color.RGBA

	// player image
	playerImg *ebiten.Image

	// player position
	playerX, playerY float64

	// player velocity
	playerVelocity float64

	// game attributes
	fullscreen  bool
	initialized bool
}

// reset game state
func (g *game) reset() {
	// center player on screen
	g.playerX = screenW/2 - playerH/2
	g.playerY = screenH/2 - playerH/2
}

func (g *game) initialize() {
	g.bgColor = color.RGBA{199, 209, 194, 255}

	// load player image
	g.playerImg = loadImage("asset/image/player.png")

	// reset game state
	g.reset()

	// initialize game attributes
	g.fullscreen = true
	g.initialized = true
}

func (g *game) Layout(outhsideWidth, outhsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func (g *game) Update() error {
	// handle initialization
	if !g.initialized {
		g.initialize()
	}

	// handle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.fullscreen = !g.fullscreen
		ebiten.SetFullscreen(g.fullscreen)
	}

	// apply gravity
	maxFallSpeed := 0.8
	if g.playerVelocity < maxFallSpeed {
		g.playerVelocity += 0.015
		if g.playerVelocity > maxFallSpeed {
			g.playerVelocity = maxFallSpeed
		}
	}

	// handle jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.playerVelocity = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if g.playerX > 0 {
			g.playerX += -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.playerX < screenW - playerW {
			g.playerX += 1
		}
	}

	// apply velocity
	g.playerY += g.playerVelocity

	// handle screen edge collision
	if g.playerY <= 0 || g.playerY >= screenH - playerH || g.playerX < 0 || g.playerX > screenW - playerW {
		g.reset()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// draw background
	screen.Fill(g.bgColor)

	// draw player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.playerImg, op)
}
