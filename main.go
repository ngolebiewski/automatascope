package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	isFullscreen bool
	toggleKey    bool // Keeps track of whether the toggle key is being held
}

func (g *Game) Update() error {
	// Check if the F key is pressed and toggle fullscreen only once per key press
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		if !g.toggleKey {
			g.isFullscreen = !g.isFullscreen
			ebiten.SetFullscreen(g.isFullscreen)
			g.toggleKey = true // Set toggleKey to true to prevent repeated toggling
		}
	} else {
		// Reset toggleKey when the key is released
		g.toggleKey = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Calculate scale factor based on screen width
	w, _ := screen.Size()
	scaleFactor := float64(w) / 320.0

	// Draw scaled text in the top-left corner of the screen
	msg := "Hello, Cellular Automatascope!\nPress 'F' to toggle fullscreen."
	ebitenutil.DebugPrintAt(screen, msg, int(10*scaleFactor), int(10*scaleFactor))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Define the game’s logical screen size
	return 320, 240
}

func main() {
	// Set initial window size and title
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Cellular Automatascope!")

	// Start the game
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
