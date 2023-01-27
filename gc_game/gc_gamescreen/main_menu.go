package gc_gamescreen

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_core/gc_ui"
)

type mainMenuScreenMode int

const (
	ScreenModeUnknown mainMenuScreenMode = iota
	ScreenModeTrademark
	ScreenModeMainMenu
)

type BuildInfo struct {
	Branch, Commit string
}

type MainMenu struct {
	screenMode mainMenuScreenMode

	renderer  gc_interface.Renderer
	uiManager *gc_ui.UIManager
	navigator gc_interface.Navigator

	buildInfo BuildInfo
}

func CreatMainMenu(navigator gc_interface.Navigator, renderer gc_interface.Renderer, ui *gc_ui.UIManager, buildInfo BuildInfo, errorMessageOptional ...string) (*MainMenu, error) {
	mainMenu := &MainMenu{
		navigator:  navigator,
		screenMode: ScreenModeUnknown,
		renderer:   renderer,
		buildInfo:  buildInfo,
		uiManager:  ui,
	}

	return mainMenu, nil
}
