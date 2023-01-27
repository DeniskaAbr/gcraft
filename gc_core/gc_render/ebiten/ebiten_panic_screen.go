package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type PanicScreen struct {
	errorMessage string
}

func (s *PanicScreen) Update() error {
	return nil
}

func (s *PanicScreen) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(colornames.Darkred)

	ebitenutil.DebugPrint(screen, s.errorMessage)
}

func (s *PanicScreen) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func CreatePanicScreen(errorMessage string) *PanicScreen {
	result := &PanicScreen{
		errorMessage: errorMessage,
	}

	ebiten.SetWindowTitle("gcraft - PANIC SCREEN")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return result
}
