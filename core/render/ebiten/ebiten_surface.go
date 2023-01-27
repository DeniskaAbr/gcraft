package ebiten

import (
	app_interface "gcraft/common/interface"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ app_interface.Surface = &ebitenSurface{}

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
