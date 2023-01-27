package gc_screen

import (
	"gcraft/gc_common/gc_interface"
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
}

func NewScreenManager(ui *gc_ui.UIManager, guiManager *gc_gui.GuiManager) *ScreenManager {
	sm := &ScreenManager{
		uiManager:  ui,
		guiManager: guiManager,
	}

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
