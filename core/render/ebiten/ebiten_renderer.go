package ebiten

import (
	"errors"
	app_interface "gcraft/common/interface"
	"gcraft/core/config"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Renderer struct {
	updateCallback
	renderCallback
	lastRenderError error
}

type renderCallback = func(surface app_interface.Surface) error

type updateCallback = func() error

func (r *Renderer) ShowPanicScreen(message string) {
	//TODO implement me
	panic("implement me")
}

const drawError = "no render callback defined for ebiten renderer"

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.lastRenderError = nil

	if r.renderCallback == nil {
		r.lastRenderError = errors.New(drawError)
		return
	}

	r.lastRenderError = r.renderCallback(r)
}

func (r *Renderer) Run(f renderCallback, u updateCallback, width, height int, title string) error {
	r.renderCallback = f
	r.updateCallback = u

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowSize(width, height)

	return ebiten.RunGame(r)
}

func (r *Renderer) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

func (r *Renderer) Update() error {
	if r.updateCallback == nil {
		return errors.New("no update callback defined for ebiten renderer")
	}

	return r.updateCallback()
}

func CreateRenderer(cfg *config.Configuration) (*Renderer, error) {
	result := &Renderer{}

	if cfg != nil {
		config := cfg

		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		ebiten.SetFullscreen(config.FullScreen)
		ebiten.SetRunnableOnUnfocused(config.RunInBackground)
		ebiten.SetVsyncEnabled(config.VsyncEnabled)
		ebiten.SetMaxTPS(config.TicksPerSecond)
	}

	return result, nil
}
