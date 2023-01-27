package ebiten

import (
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
