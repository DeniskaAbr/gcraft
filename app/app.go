package app

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"gcraft/gc_core/gc_screen"
	"gcraft/gc_core/gc_ui"

	"gcraft/gc_core/gc_gui"

	"golang.org/x/image/colornames"

	"gcraft/gc_common/gc_interface"
	"gcraft/gc_core/gc_config"
	"gcraft/gc_core/gc_render/ebiten"
)

type App struct {
	lastTime          float64
	lastScreenAdvance float64
	timeScale         float64
	guiManager        *gc_gui.GuiManager
	renderer          gc_interface.Renderer
	screen            *gc_screen.ScreenManager
	ui                *gc_ui.UIManager
	errorMessage      error
	config            *gc_config.Configuration
	gitBranch         string
	gitCommit         string
	*Options
}

type Options struct{}

const (
	errMsgPadding = 20
)

const (
	appLoggerPrefix = "App"
)

func Create(gitBranch, gitCommit string) *App {
	runtime.LockOSThread()

	app := &App{
		gitBranch: gitBranch,
		gitCommit: gitCommit,
		Options:   &Options{},
	}

	app.parseArguments()

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

	windowTitle := fmt.Sprintf("gocraft (%s)", a.gitBranch)

	if err := a.initialize(); err != nil {
		if a.errorMessage == nil {
			a.errorMessage = err
		}

		gameErr := a.renderer.Run(a.updateInitError, updateNOOP, 800, 600, windowTitle)
		if gameErr != nil {
			return gameErr
		}
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
	// a.errorMessage = errors.New("asd")

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

func (a *App) parseArguments() {
	showVersion := flag.Bool("v", false, "Show version")
	showHelp := flag.Bool("h", false, "Show help")

	flag.Usage = func() {
		fmt.Printf("usage: %s [<flags>]\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		// a.Infof("version: OpenDiablo2 (%s %s)", a.gitBranch, a.gitCommit)
		os.Exit(0)
	}

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}
}
