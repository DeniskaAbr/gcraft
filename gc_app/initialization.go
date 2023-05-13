package gc_app

import (
	"gcraft/gc_common/gc_util"
	"gcraft/gc_core/gc_config"
	"gcraft/gc_core/gc_gui"
	"gcraft/gc_core/gc_screen"
)

func (a *App) initialize() error {
	if err := a.initConfig(a.config); err != nil {
		return err
	}

	a.timeScale = 1.0
	a.lastTime = gc_util.Now()
	a.lastScreenAdvance = a.lastTime

	a.renderer.SetWindowIcon("logo.png")

	gui, err := gc_gui.CreateGuiManager(*a.Options.LogLevel, a.inputManager)
	if err != nil {
		return err
	}

	a.guiManager = gui

	a.screen = gc_screen.NewScreenManager(a.ui, *a.Options.LogLevel, a.guiManager)

	// a.audio.SetVolumes()

	a.ui.Initialize()

	return nil
}

func (a *App) initConfig(config *gc_config.Configuration) error {
	a.config = config

	return nil
}
