package ebiten

import (
	"fmt"
	"image/color"

	"gcraft/gc_common/gc_interface"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ gc_interface.Surface = &ebitenSurface{}

type ebitenSurface struct {
	renderer     *Renderer
	stateStack   []surfaceState
	stateCurrent surfaceState
	image        *ebiten.Image
}

func (s *ebitenSurface) GetDepth() int {
	return len(s.stateStack)
}

func createEbitenSurface(r *Renderer, img *ebiten.Image, currentState ...surfaceState) *ebitenSurface {
	state := surfaceState{}

	if len(currentState) > 0 {
		state = currentState[0]
	}

	return &ebitenSurface{
		renderer:     r,
		image:        img,
		stateCurrent: state,
	}
}

func (s *ebitenSurface) Renderer() gc_interface.Renderer {
	return s.renderer
}

func (s *ebitenSurface) Clear(fillColor color.Color) {
	s.image.Fill(fillColor)
}

func (s *ebitenSurface) PushTranslation(x, y int) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

func (s *ebitenSurface) DrawTextf(format string, params ...interface{}) {
	str := fmt.Sprintf(format, params...)
	s.Renderer().PrintAt(s.image, str, s.stateCurrent.x, s.stateCurrent.y)
}
