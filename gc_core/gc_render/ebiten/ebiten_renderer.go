package ebiten

import (
	"errors"
	"image"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"gcraft/gc_common/gc_util"

	"gcraft/gc_common/gc_interface"
	"gcraft/gc_core/gc_config"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Renderer struct {
	updateCallback
	renderCallback
	*gc_util.GlyphPrinter
	lastRenderError error
}

type renderCallback = func(surface gc_interface.Surface) error

type updateCallback = func() error

var _ gc_interface.Renderer = &Renderer{}

func (r *Renderer) ShowPanicScreen(message string) {
	errorScreen := CreatePanicScreen(message)

	err := ebiten.RunGame(errorScreen)
	if err != nil {
		panic(err)
	}
}

const drawError = "no render callback defined for ebiten renderer"

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.lastRenderError = nil

	if r.renderCallback == nil {
		r.lastRenderError = errors.New(drawError)
		return
	}

	r.lastRenderError = r.renderCallback(createEbitenSurface(r, screen))
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

func CreateRenderer(cfg *gc_config.Configuration) (*Renderer, error) {
	result := &Renderer{
		GlyphPrinter: gc_util.NewDebugPrinter(),
	}

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

func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}
