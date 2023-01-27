package gc_screen

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_util"
	"gcraft/gc_core/gc_gui"
	"gcraft/gc_core/gc_ui"
)

const (
	logPrefix = "Screen Manager"
)

type ScreenManager struct {
	nextScreen Screen
	uiManager  *gc_ui.UIManager
	guiManager *gc_gui.GuiManager

	*gc_util.Logger
}

func NewScreenManager(ui *gc_ui.UIManager, l gc_util.LogLevel, guiManager *gc_gui.GuiManager) *ScreenManager {
	sm := &ScreenManager{
		uiManager:  ui,
		guiManager: guiManager,
	}

	sm.Logger = gc_util.NewLogger()
	sm.Logger.SetPrefix(logPrefix)
	sm.Logger.SetLevel(l)

	return sm
}

func (sm *ScreenManager) Advance(elapsed float64) error {
	switch {
	}

	return nil
}

func (sm *ScreenManager) SetNextScreen(screen Screen) {
	sm.nextScreen = screen
}

func (sm *ScreenManager) Render(surface gc_interface.Surface) {
}
