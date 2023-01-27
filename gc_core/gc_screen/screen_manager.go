package gc_screen

import (
	"gcraft/gc_core/gc_gui"
	"gcraft/gc_core/gc_ui"
)

const (
	logPrefix = "Screen Manager"
)

type ScreenManager struct {
	uiManager  *gc_ui.UIManager
	guiManager *gc_gui.GuiManager
}

func NewScreenManager(ui *gc_ui.UIManager, guiManager *gc_gui.GuiManager) *ScreenManager {
	sm := &ScreenManager{
		uiManager:  ui,
		guiManager: guiManager,
	}

	return sm
}
