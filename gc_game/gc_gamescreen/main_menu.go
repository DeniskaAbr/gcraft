package gc_gamescreen

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_util"
	"gcraft/gc_core/gc_ui"
)

type mainMenuScreenMode int

const (
	ScreenModeUnknown mainMenuScreenMode = iota
	ScreenModeTrademark
	ScreenModeMainMenu
)

const (
	logPrefix = "Game Screen"
)

type BuildInfo struct {
	Branch, Commit string
}

type MainMenu struct {
	screenMode mainMenuScreenMode

	renderer  gc_interface.Renderer
	uiManager *gc_ui.UIManager
	navigator gc_interface.Navigator

	audioProvider gc_interface.AudioProvider

	buildInfo BuildInfo

	*gc_util.Logger
}

func CreatMainMenu(
	navigator gc_interface.Navigator,
	renderer gc_interface.Renderer,
	audioProvider gc_interface.AudioProvider,
	ui *gc_ui.UIManager,
	buildInfo BuildInfo,
	l gc_util.LogLevel,
	errorMessageOptional ...string,
) (*MainMenu, error) {
	mainMenu := &MainMenu{
		navigator:     navigator,
		screenMode:    ScreenModeUnknown,
		renderer:      renderer,
		audioProvider: audioProvider,
		buildInfo:     buildInfo,
		uiManager:     ui,
	}

	mainMenu.Logger = gc_util.NewLogger()
	mainMenu.Logger.SetPrefix(logPrefix)
	mainMenu.Logger.SetLevel(l)

	if len(errorMessageOptional) != 0 {
	}

	return mainMenu, nil
}
