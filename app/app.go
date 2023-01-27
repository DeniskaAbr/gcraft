package app

import (
	app_interface "gcraft/common/interface"
	"gcraft/core/config"
	"gcraft/core/render/ebiten"
	"runtime"
)

type App struct {
	renderer     app_interface.Renderer
	errorMessage error
	config       *config.Configuration
	*Options
}

type Options struct {
}

const ()

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
	if a.errorMessage != nil {
		return a.renderer.Run(a.updateInitError, updateNOOP, 800, 600, "gcraft")
	}

	return nil
}

func updateNOOP() error {

	return nil
}

func (a *App) updateInitError(target app_interface.Surface) error {

	return nil
}
