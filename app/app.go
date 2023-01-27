package app

import (
	"errors"
	"runtime"

	"golang.org/x/image/colornames"

	"gcraft/gc_common/gc_interface"
	"gcraft/gc_core/gc_config"
	"gcraft/gc_core/gc_render/ebiten"
)

type App struct {
	renderer     gc_interface.Renderer
	errorMessage error
	config       *gc_config.Configuration
	*Options
}

type Options struct{}

const (
	errMsgPadding = 20
)

const (
	appLoggerPrefix = "App"
)

func Create() *App {
	runtime.LockOSThread()

	app := &App{
		Options: &Options{},
	}

	return app
}

func (a *App) Run() (err error) {
	if a.config, err = a.LoadConfig(); err != nil {
		return err
	}

	if err := a.loadEngine(); err != nil {
		a.renderer.ShowPanicScreen(err.Error())
		return err
	}

	return nil
}

func (a *App) loadEngine() error {
	renderer, err := ebiten.CreateRenderer(a.config)
	if err != nil {
		return err
	}

	a.renderer = renderer

	// test error
	a.errorMessage = errors.New("asd")

	if a.errorMessage != nil {
		return a.renderer.Run(a.updateInitError, updateNOOP, 800, 600, "gcraft")
	}

	return nil
}

func updateNOOP() error {
	return nil
}

func (a *App) updateInitError(target gc_interface.Surface) error {
	target.Clear(colornames.Darkred)
	target.PushTranslation(errMsgPadding, errMsgPadding)
	target.DrawTextf(a.errorMessage.Error())

	return nil
}

func (a *App) LoadConfig() (*gc_config.Configuration, error) {
	config := &gc_config.Configuration{}
	return config, nil
}
